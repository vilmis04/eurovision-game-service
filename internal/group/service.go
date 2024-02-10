package group

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/vilmis04/eurovision-game-service/internal/utils"
)

type Service struct {
	storage *Repo
}

func NewService() *Service {
	return &Service{
		storage: NewRepo(),
	}
}

func (s *Service) GetGroups(owner string, request *http.Request) (*[]byte, error) {
	groupName := request.URL.Query().Get("name")
	groups, err := s.storage.GetGroupList(owner, groupName)
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
	err := utils.DecodeRequestJson(request, requestBody)
	if err != nil {
		return nil, err
	}

	name := *requestBody.Name
	usedNames, err := s.storage.GetGroupNames(owner)
	if err != nil {
		return nil, err
	}

	if slices.Contains(*usedNames, name) {
		return nil, fmt.Errorf("group %v already exists", name)
	}

	group := Group{
		Name:        name,
		Owner:       owner,
		DateCreated: time.Now().String(),
		Members:     []string{owner},
	}

	id, err := s.storage.CreateGroup(&group)
	if err != nil {
		return nil, err
	}

	encodedId, err := json.Marshal(id)
	if err != nil {
		return nil, err
	}

	return &encodedId, nil
}
