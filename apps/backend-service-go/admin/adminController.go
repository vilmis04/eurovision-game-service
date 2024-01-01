package admin

import "fmt"

type adminController struct{}

func NewController() *adminController {
	return &adminController{}
}

func (ctrl *adminController) GetCofig() (*AdminConfig, error) {
	// TODO: add get config
	fmt.Printf("Get Config")
	var err error = nil

	return &AdminConfig{}, err
}

func (ctrl *adminController) UpdateConfig() error {
	// TODO: add update config
	fmt.Printf("Update Config")

	return nil
}

func (ctrl *adminController) GetCountryDetails() {

}

func (ctrl *adminController) UpdateCountryDetails() {

}

func (ctrl *adminController) CreateCountryDetails() {

}

func (ctrl *adminController) DeleteCountry() {

}
