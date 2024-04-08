package group

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"
)

type Service struct {
	Repo
}

func NewService() *Service {
	return &Service{
		Repo: *NewRepo(),
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

// TODO: add check that only one group with the same name exists per user (owner)
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

func (s *Service) DeleteGroup(owner string, name string) error {
	err := s.Repo.DeleteGroup(owner, name)
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
	updatedMemberList := append(group.Members, user)

	err = s.Repo.UpdateMembers(groupOwner, groupName, updatedMemberList)
	if err != nil {
		return err
	}

	return nil
}
