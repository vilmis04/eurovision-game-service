package country

import "fmt"

type countryService struct{}

func NewService() *countryService {
	return &countryService{}
}

func (s *countryService) GetCountry() (*Country, error) {
	// TODO: add get country details
	fmt.Println("Get country details")

	return &Country{}, nil
}

func (s *countryService) UpdateCountry() error {
	// TODO: Take an object, update accordingly and return error
	fmt.Println("Update country details")

	return nil
}

func (s *countryService) CreateCountry() error {
	// TODO: Take an object, create and return error
	fmt.Println("Create details for new country")

	return nil
}

func (s *countryService) DeleteCountry() error {
	// TODO: Take an object, create and return error
	fmt.Println("Delete by country details by id")

	return nil
}
