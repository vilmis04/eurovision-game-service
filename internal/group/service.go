package group

import (
	"cmp"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/lib/pq"
	"github.com/vilmis04/eurovision-game-service/internal/score"
	"github.com/vilmis04/eurovision-game-service/internal/utils"
)

// store is the persistence the service needs, implemented by Repo.
type store interface {
	GetGroupList(user string, groupId string) (*[]Group, error)
	GetGroupById(id int64) (*Group, error)
	GetGroupNames(owner string) (*[]string, error)
	CreateGroup(group *Group) (*int64, error)
	UpdateMembers(id int64, members []string) error
	DeleteGroup(owner string, id int64) error
}

const (
	maxGroupNameLength  = 20  // "group"."name" is VARCHAR(20)
	maxMemberLength     = 255 // matches "score"."user"
	maxMembersPerUpdate = 100
)

type Service struct {
	store        store
	scoreService *score.Service
	inviteSecret []byte
	now          func() time.Time
}

func NewService(db *sql.DB, scoreService *score.Service) *Service {
	return &Service{
		store:        NewRepo(db),
		scoreService: scoreService,
		inviteSecret: []byte(os.Getenv("INVITE_SECRET")),
		now:          time.Now,
	}
}

// parseGroupId validates a user supplied group id. An empty id is allowed
// only when allowEmpty is set and yields 0.
func parseGroupId(id string, allowEmpty bool) (int64, error) {
	if id == "" && allowEmpty {
		return 0, nil
	}
	parsed, err := strconv.ParseInt(id, 10, 64)
	if err != nil || parsed <= 0 {
		return 0, utils.BadRequest("invalid group id")
	}

	return parsed, nil
}

func (s *Service) GetGroups(user string, request *http.Request) (*[]byte, error) {
	groupId := request.URL.Query().Get("id")
	if _, err := parseGroupId(groupId, true); err != nil {
		return nil, err
	}

	groups, err := s.store.GetGroupList(user, groupId)
	if err != nil {
		return nil, err
	}

	encodedGroups, err := json.Marshal(groups)
	if err != nil {
		return nil, err
	}

	return &encodedGroups, nil
}

func (s *Service) CreateGroup(owner string, request *http.Request) (*[]byte, error) {
	var requestBody CreateGroupRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		return nil, utils.BadRequest("invalid request body")
	}

	name := strings.TrimSpace(requestBody.Name)
	if name == "" || utf8.RuneCountInString(name) > maxGroupNameLength {
		return nil, utils.BadRequest("invalid group name")
	}

	usedNames, err := s.store.GetGroupNames(owner)
	if err != nil {
		return nil, err
	}

	if slices.Contains(*usedNames, name) {
		return nil, utils.Conflict("group already exists")
	}

	group := Group{
		Name:        name,
		Owner:       owner,
		DateCreated: s.now(),
		Members:     []string{owner},
	}

	id, err := s.store.CreateGroup(&group)
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" { // unique_violation: lost a race with a same-named group
		return nil, utils.Conflict("group already exists")
	}
	if err != nil {
		return nil, err
	}

	encodedId, err := json.Marshal(id)
	if err != nil {
		return nil, err
	}

	return &encodedId, nil
}

// ownedGroup loads the group and verifies that the user is its owner.
func (s *Service) ownedGroup(user string, rawGroupId string) (*Group, error) {
	groupId, err := parseGroupId(rawGroupId, false)
	if err != nil {
		return nil, err
	}

	group, err := s.store.GetGroupById(groupId)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, utils.NotFound("group not found")
	}
	if group.Owner != user {
		return nil, utils.Forbidden("only the group owner can do this")
	}

	return group, nil
}

func (s *Service) UpdateMembers(user string, rawGroupId string, request *http.Request) error {
	group, err := s.ownedGroup(user, rawGroupId)
	if err != nil {
		return err
	}

	var requestBody UpdateGroupRequestBody
	err = json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		return utils.BadRequest("invalid request body")
	}
	if len(requestBody.Members) > maxMembersPerUpdate {
		return utils.BadRequest("too many members")
	}

	updatedMemberList := slices.Clone(group.Members)
	for _, member := range requestBody.Members {
		member = strings.TrimSpace(member)
		if member == "" || utf8.RuneCountInString(member) > maxMemberLength {
			return utils.BadRequest("invalid member")
		}
		if !slices.Contains(updatedMemberList, member) {
			updatedMemberList = append(updatedMemberList, member)
		}
	}

	return s.store.UpdateMembers(group.Id, updatedMemberList)
}

func (s *Service) DeleteGroup(owner string, rawGroupId string) error {
	groupId, err := parseGroupId(rawGroupId, false)
	if err != nil {
		return err
	}

	return s.store.DeleteGroup(owner, groupId)
}

func (s *Service) GenerateInvite(rawGroupId string, user string) (string, error) {
	group, err := s.ownedGroup(user, rawGroupId)
	if err != nil {
		return "", err
	}

	return createInvite(s.inviteSecret, group.Id, s.now().Add(InviteTTL))
}

func (s *Service) JoinGroup(user string, request *http.Request) error {
	var requestBody JoinGroupRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		return utils.BadRequest("invalid request body")
	}

	groupId, err := parseInvite(s.inviteSecret, requestBody.InviteCode, s.now())
	if errors.Is(err, errInviteNoSecret) {
		return err
	}
	if err != nil {
		return utils.BadRequest("invalid or expired invite")
	}

	group, err := s.store.GetGroupById(groupId)
	if err != nil {
		return err
	}
	if group == nil {
		return utils.NotFound("group not found")
	}
	if slices.Contains(group.Members, user) {
		return nil
	}

	return s.store.UpdateMembers(group.Id, append(slices.Clone(group.Members), user))
}

func (s *Service) getGroupList(user string, groupId string) (map[int64]string, *[]Group, error) {
	if _, err := parseGroupId(groupId, true); err != nil {
		return nil, nil, err
	}

	allGroupMap := make(map[int64]string)
	allGroups, err := s.store.GetGroupList(user, "")
	if err != nil {
		return nil, nil, err
	}
	for _, group := range *allGroups {
		allGroupMap[group.Id] = group.Name
	}

	groupList, err := s.store.GetGroupList(user, groupId)
	if err != nil {
		return nil, nil, err
	}

	return allGroupMap, groupList, nil
}

func (s *Service) getPlayerResults(groupList *[]Group) (map[string]Member, error) {
	memberMap := make(map[string]Member)
	memberList := []string{}
	for _, group := range *groupList {
		for _, member := range group.Members {
			_, ok := memberMap[member]
			if !ok {
				memberList = append(memberList, member)
				memberMap[member] = Member{
					Name: member,
				}
			}
		}
	}

	pointsMap, err := s.scoreService.GetMultipleUserScores(memberList)
	if err != nil {
		return nil, err
	}
	for member, results := range memberMap {
		points := pointsMap[member]
		memberMap[member] = Member{
			Name:  results.Name,
			Score: points,
		}
	}

	return memberMap, nil
}

func (s *Service) getSortedPoints(memberMap map[string]Member) []uint16 {
	pointList := []uint16{}
	for _, member := range memberMap {
		if slices.Contains(pointList, member.Score) {
			continue
		}
		pointList = append(pointList, member.Score)
	}
	slices.SortStableFunc(pointList, func(a, b uint16) int { return cmp.Compare(b, a) })

	return pointList
}

func (s *Service) getSortedPlayerResults(groupList *[]Group) ([]Member, error) {
	memberMap, err := s.getPlayerResults(groupList)
	if err != nil {
		return nil, err
	}

	pointList := s.getSortedPoints(memberMap)

	memberList := []Member{}
	for _, member := range memberMap {
		memberList = append(memberList, Member{
			Name:     member.Name,
			Score:    member.Score,
			Position: slices.Index(pointList, member.Score) + 1,
		})
	}
	slices.SortStableFunc(memberList, func(a, b Member) int { return cmp.Compare(a.Position, b.Position) })

	return memberList, nil
}

func (s *Service) GetLeaderboard(user string, groupId string) (*[]byte, error) {
	allGroupList, groupList, err := s.getGroupList(user, groupId)
	if err != nil {
		return nil, err
	}

	memberList, err := s.getSortedPlayerResults(groupList)
	if err != nil {
		return nil, err
	}

	leaderboard := Leaderboard{
		Groups:     allGroupList,
		PlayerList: memberList,
	}
	response, err := json.Marshal(leaderboard)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
