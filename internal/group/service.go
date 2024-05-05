package group

import (
	"cmp"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/vilmis04/eurovision-game-service/internal/score"
)

type Service struct {
	Repo
	scoreService score.Service
}

func NewService() *Service {
	return &Service{
		Repo:         *NewRepo(),
		scoreService: *score.NewService(),
	}
}

func (s *Service) GetGroups(user string, request *http.Request) (*[]byte, error) {
	groupId := request.URL.Query().Get("id")
	groups, err := s.Repo.GetGroupList(user, groupId)
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
		return nil, fmt.Errorf("service: %v", err)
	}

	name := requestBody.Name
	usedNames, err := s.Repo.GetGroupNames(owner)
	if err != nil {
		return nil, err
	}

	if slices.Contains(*usedNames, name) {
		return nil, fmt.Errorf("group %v already exists", name)
	}

	group := Group{
		Name:        name,
		Owner:       owner,
		DateCreated: time.Now(),
		Members:     []string{owner},
	}

	id, err := s.Repo.CreateGroup(&group)
	if err != nil {
		return nil, err
	}

	encodedId, err := json.Marshal(id)
	if err != nil {
		return nil, err
	}

	return &encodedId, nil
}

func (s *Service) UpdateMembers(owner string, name string, request *http.Request) error {
	if owner == "" || name == "" {
		return fmt.Errorf("%v", http.StatusBadRequest)
	}
	var requestBody UpdateGroupRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		return err
	}

	groupList, err := s.Repo.GetGroupList(owner, name)
	if err != nil {
		return err
	}

	numberOfGroups := len(*groupList)
	if numberOfGroups == 0 {
		return fmt.Errorf("no group named %v found", name)
	}
	if numberOfGroups > 1 {
		return fmt.Errorf("multiple groups named %v found", name)
	}
	group := (*groupList)[0]

	updatedMemberList := append(group.Members, requestBody.Members...)

	return s.Repo.UpdateMembers(owner, name, updatedMemberList)
}

func (s *Service) DeleteGroup(owner string, id string) error {
	err := s.Repo.DeleteGroup(owner, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GenerateInvite(id string, user string) (string, error) {
	groupList, err := s.GetGroupList(user, id)
	if err != nil {
		return "", err
	}
	name := (*groupList)[0].Name

	date := time.Now().Format("2006-01-02") // 2006-01-02 directs the format in Go for YYYY-DD-MM
	message := fmt.Sprintf("%v:%v:%v:%v", name, user, id, date)
	link := base64.RawStdEncoding.EncodeToString([]byte(message))

	return link, nil
}

func (s *Service) JoinGroup(user string, request *http.Request) error {
	var requestBody JoinGroupRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		return fmt.Errorf("service: %v", err)
	}

	link, err := base64.RawStdEncoding.DecodeString(requestBody.InviteCode)
	if err != nil {
		return fmt.Errorf("service: %v", err)
	}

	inviteInfo := strings.Split(string(link), ":")
	groupName := inviteInfo[0]
	groupOwner := inviteInfo[1]
	groupId := inviteInfo[2]

	groupList, err := s.Repo.GetGroupList(groupOwner, groupId)
	if err != nil {
		return err
	}
	group := (*groupList)[0]
	updatedMemberList := group.Members
	if !slices.Contains(updatedMemberList, user) {
		updatedMemberList = append(group.Members, user)
	}

	err = s.Repo.UpdateMembers(groupOwner, groupName, updatedMemberList)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) getGroupList(user string, groupId string) (map[int64]string, *[]Group, error) {

	allGroupList := make(map[int64]string)
	allGroups, err := s.GetGroupList(user, "")
	if err != nil {
		return nil, nil, err
	}
	for _, group := range *allGroups {
		allGroupList[group.Id] = group.Name
	}

	groupList, err := s.GetGroupList(user, groupId)
	if err != nil {
		return nil, nil, err
	}

	return allGroupList, groupList, nil
}

func (s *Service) getPlayerResults(groupList *[]Group) (map[string]Member, error) {
	memberMap := make(map[string]Member)
	for _, group := range *groupList {
		for _, member := range group.Members {
			_, ok := memberMap[member]
			if !ok {
				points, err := s.scoreService.GetUserScore(member)
				if err != nil {
					return nil, err
				}
				memberMap[member] = Member{
					Name:  member,
					Score: points,
				}
			}
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
