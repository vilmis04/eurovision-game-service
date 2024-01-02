package admin

import (
	"fmt"
)

type adminService struct{}

func NewService() *adminService {
	return &adminService{}
}

func (s *adminService) GetConfig() (*AdminConfig, error) {
	// TODO: add get config
	fmt.Println("Get Config")

	return &AdminConfig{}, nil
}

func (s *adminService) UpdateConfig() error {
	// TODO: add update config
	fmt.Println("Update Config")

	return nil
}

func (s *adminService) GetCountryDetails() (*CountryDetails, error) {
	// TODO: add get country details
	fmt.Println("Get country details")

	return &CountryDetails{}, nil
}

func (s *adminService) UpdateCountryDetails() error {
	// TODO: Take an object, update accordingly and return error
	fmt.Println("Update country details")

	return nil
}

func (s *adminService) CreateCountryDetails() error {
	// TODO: Take an object, create and return error
	fmt.Println("Create details for new country")

	return nil
}

func (s *adminService) DeleteCountry() error {
	// TODO: Take an object, create and return error
	fmt.Println("Delete by country details by id")

	return nil
}
