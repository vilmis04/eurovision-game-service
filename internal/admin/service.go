package admin

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
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

func (s *Service) GetConfig() (*[]byte, error) {

	config, err := s.storage.GetConfig()
	if err != nil {
		return nil, err
	}

	encodedConfig, err := json.Marshal(&config)
	if err != nil {
		return nil, err
	}

	return &encodedConfig, nil
}

func (s *Service) UpdateConfig(req *http.Request) error {
	var body adminConfigRequestBody

	err := utils.DecodeRequestJson(req, &body)
	if err != nil {
		return err
	}

	return s.storage.UpdateConfig(&body)
}
