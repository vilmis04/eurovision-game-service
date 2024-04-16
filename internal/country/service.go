package country

import (
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

func NewService() *Service {
	return &Service{
		storage:      NewRepo(),
		adminService: admin.NewService(),
	}
}

func (s *Service) CreateCountry(request *http.Request) (*[]byte, error) {
	var requestBody CreateCountryRequest
	err := utils.DecodeRequestJson(request, &requestBody)
	if err != nil {
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

func (s *Service) GetCountryList(year string, request *http.Request) (*[]byte, error) {
	queryParams := request.URL.Query()
	gameType := queryParams.Get("gameType")
	name := queryParams.Get("name")

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
	var requestBody UpdateCountryRequest
	err := utils.DecodeRequestJson(request, requestBody)
	if err != nil {
		return err
	}

	countryList, err := s.storage.GetCountryList(params["year"], "", params["name"])
	if err != nil {
		return err
	}
	if len(*countryList) > 1 {
		return fmt.Errorf("query returned multiple countries: %v", http.StatusBadRequest)
	}
	if len(*countryList) == 0 {
		return fmt.Errorf("country not found: %v", http.StatusNotFound)
	}

	return s.storage.UpdateCountry(&requestBody, &params)
}

func (s *Service) DeleteCountry(params *map[string]string) error {
	return s.storage.DeleteCountry((*params)["year"], (*params)["name"])
}
