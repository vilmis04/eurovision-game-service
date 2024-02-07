package group

import (
	"encoding/json"
	"net/http"
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
