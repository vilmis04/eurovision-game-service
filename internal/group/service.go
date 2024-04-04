package group

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
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

func (s *Service) GetGroups(owner string, request *http.Request) (*[]byte, error) {
	groupName := request.URL.Query().Get("name")
	groups, err := s.Repo.GetGroupList(owner, groupName)
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

func (s *Service) UpdateMembers(owner string, request *http.Request) error {
	var requestBody UpdateGroupRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		return err
	}

	groupList, err := s.Repo.GetGroupList(owner, *requestBody.Name)
	if err != nil {
		return err
	}

	numberOfGroups := len(*groupList)
	if numberOfGroups == 0 {
		return fmt.Errorf("no group named %v found", requestBody.Name)
	}
	if numberOfGroups > 1 {
		return fmt.Errorf("multiple groups named %v found", requestBody.Name)
	}
	group := (*groupList)[0]

	return s.Repo.UpdateGroup(owner, &requestBody, &group)
}

func (s *Service) DeleteGroup(owner string, name string) error {
	err := s.Repo.DeleteGroup(owner, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GenerateInvite(name string, user string) (string, error) {
	date := time.Now().Format("2006-01-02") // 2006-01-02 directs the format in Go for YYYY-DD-MM
	message := fmt.Sprintf("%v:%v:%v", name, user, date)
	link := base64.RawStdEncoding.EncodeToString([]byte(message))

	return link, nil
}
