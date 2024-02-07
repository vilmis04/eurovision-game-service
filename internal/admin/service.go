package admin

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vilmis04/eurovision-game-service/internal/utils"
)

type AdminService struct {
	storage *AdminRepo
}

func NewService() *AdminService {
	return &AdminService{
		storage: NewRepo(),
	}
}

func (s *AdminService) GetConfig() (*[]byte, error) {

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

func (s *AdminService) UpdateConfig(req *http.Request) error {
	var body adminConfigRequestBody

	err := utils.DecodeRequestJson(req, &body)
	if err != nil {
		return err
	}

	return s.storage.UpdateConfig(&body)
}
