package country

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/utils"
)

type Service struct {
	storage      *Repo
	adminService *admin.Service
}

func NewService(db *sql.DB, adminService *admin.Service) *Service {
	return &Service{
		storage:      NewRepo(db),
		adminService: adminService,
	}
}

func (s *Service) CreateCountry(request *http.Request) (*[]byte, error) {
	var requestBody CreateCountryRequest
	err := utils.DecodeRequestJson(request, &requestBody)
	if err != nil {
		return nil, utils.BadRequest("invalid request body")
	}
	if err := requestBody.validate(); err != nil {
		return nil, err
	}

	encodedConfig, err := s.adminService.GetConfig()
	if err != nil {
		return nil, err
	}

	var config admin.Admin
	err = json.Unmarshal(*encodedConfig, &config)
	if err != nil {
		return nil, err
	}

	country := Country{
		Year:       config.Year,
		IsInFinal:  false,
		Score:      0,
		Name:       *requestBody.Name,
		Code:       *requestBody.Code,
		GameType:   *requestBody.GameType,
		Artist:     *requestBody.Artist,
		Song:       *requestBody.Song,
		OrderSemi:  *requestBody.OrderSemi,
		OrderFinal: 0,
	}

	id, err := s.storage.Create(&country)
	if err != nil {
		return nil, err
	}

	encodedId, err := json.Marshal(id)
	if err != nil {
		return nil, err
	}

	return &encodedId, nil
}

func (s *Service) GetCountrySummary(year string, gameType string, name string) (*[]byte, error) {
	if _, err := parseYear(year); err != nil {
		return nil, err
	}

	countries, err := s.storage.GetCountrySummary(year, gameType, name)
	if err != nil {
		return nil, err
	}

	encodedCountries, err := json.Marshal(countries)
	if err != nil {
		return nil, err
	}
	return &encodedCountries, nil
}

func (s *Service) GetCountryList(year string, gameType string, name string) (*[]byte, error) {
	countries, err := s.storage.GetCountryList(year, gameType, name)
	if err != nil {
		return nil, err
	}

	encodedCountries, err := json.Marshal(countries)
	if err != nil {
		return nil, err
	}
	return &encodedCountries, nil
}

func (s *Service) UpdateCountry(params map[string]string, request *http.Request) error {
	year, err := parseYear(params["year"])
	if err != nil {
		return err
	}
	name, err := parseName(params["name"])
	if err != nil {
		return err
	}

	var requestBody UpdateCountryRequest
	err = utils.DecodeRequestJson(request, &requestBody)
	if err != nil {
		return utils.BadRequest("invalid request body")
	}
	if requestBody.Score == nil && requestBody.IsInFinal == nil && requestBody.OrderSemi == nil && requestBody.OrderFinal == nil {
		return utils.BadRequest("nothing to update")
	}

	updated, err := s.storage.UpdateCountry(&requestBody, year, name)
	if err != nil {
		return err
	}
	if updated == 0 {
		return utils.NotFound("country not found")
	}

	return nil
}

func (s *Service) DeleteCountry(params *map[string]string) error {
	year, err := parseYear((*params)["year"])
	if err != nil {
		return err
	}
	name, err := parseName((*params)["name"])
	if err != nil {
		return err
	}

	deleted, err := s.storage.DeleteCountry(year, name)
	if err != nil {
		return err
	}
	if deleted == 0 {
		return utils.NotFound("country not found")
	}
	if deleted > 1 {
		return fmt.Errorf("delete country %v %v matched %v rows, nothing deleted", year, name, deleted)
	}

	return nil
}

func (s *Service) GetFinalists(year uint16) ([]Country, error) {
	finalists, err := s.storage.GetFinalists(year)
	if err != nil {
		return nil, err
	}

	return finalists, nil
}
