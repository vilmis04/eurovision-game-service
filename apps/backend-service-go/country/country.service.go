package country

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/admin"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/utils"
)

type CountryService struct {
	storage      *CountryRepo
	adminService *admin.AdminService
}

func NewService() *CountryService {
	return &CountryService{
		storage:      NewRepo(),
		adminService: admin.NewService(),
	}
}

func (s *CountryService) CreateCountry(request *http.Request) (*[]byte, error) {
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
		Year:      config.Year,
		IsInFinal: false,
		Score:     0,
		Name:      *requestBody.Name,
		GameType:  *requestBody.GameType,
		Artist:    *requestBody.Artist,
		Song:      *requestBody.Song,
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

func (s *CountryService) GetAllCountries() (*[]Country, error) {
	// TODO: add get country details
	fmt.Println("Get country details")

	return &[]Country{}, nil
}

func (s *CountryService) GetCountry() (*Country, error) {
	// TODO: add get country details
	fmt.Println("Get country details")

	return &Country{}, nil
}

func (s *CountryService) UpdateCountry() error {
	// TODO: Take an object, update accordingly and return error
	fmt.Println("Update country details")

	return nil
}

func (s *CountryService) DeleteCountry() error {
	// TODO: Take an object, create and return error
	fmt.Println("Delete by country details by id")

	return nil
}
