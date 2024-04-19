package score

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/country"
)

type Service struct {
	storage        *Repo
	adminService   *admin.Service
	countryService *country.Service
}

func NewService() *Service {
	return &Service{
		storage:        NewRepo(),
		adminService:   admin.NewService(),
		countryService: country.NewService(),
	}
}

func (s *Service) InitializeScores(user string, year uint16, gameType admin.GameType) ([]ScoreResponse, error) {
	encodedCountries, err := s.countryService.GetCountryList(fmt.Sprint(year), fmt.Sprint(gameType), "")
	if err != nil {
		return nil, err
	}
	var countries []country.Country
	err = json.Unmarshal(*encodedCountries, &countries)
	if err != nil {
		return nil, err
	}

	scores, err := s.storage.InitializeScores(user, year, gameType, &countries)
	if err != nil {
		return nil, err
	}

	return scores, nil
}

func (s *Service) UpdateScore(country string, user string, request *http.Request) error {
	encodedConfig, err := s.adminService.GetConfig()
	if err != nil {
		return err
	}
	var config admin.Admin
	err = json.Unmarshal(*encodedConfig, &config)
	if err != nil {
		return err
	}

	// TODO: add update logic

	return nil
}

func (s *Service) GetAllScores(user string) (*[]byte, error) {
	encodedConfig, err := s.adminService.GetConfig()
	if err != nil {
		return nil, err
	}
	var config admin.Admin
	err = json.Unmarshal(*encodedConfig, &config)
	if err != nil {
		return nil, err
	}

	scores, err := s.storage.GetAllScores(user, config.GameType, config.Year)
	if err != nil {
		return nil, fmt.Errorf("failed to get scores for user %s: %v", user, err)
	}
	if len(scores) == 0 {
		scores, err = s.InitializeScores(user, config.Year, config.GameType)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize scores for user %s: %v", user, err)
		}
	}

	encodedScores, err := json.Marshal(scores)
	if err != nil {
		return nil, err
	}

	return &encodedScores, nil
}
