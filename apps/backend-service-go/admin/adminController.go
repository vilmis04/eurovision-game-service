package admin

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type adminController struct {
	service *adminService
}

func NewController() *adminController {
	return &adminController{
		service: NewService(),
	}
}

func (ctrl *adminController) GetConfig(c *gin.Context) {
	config, err := ctrl.service.GetConfig()
	if err != nil {
		fmt.Printf("Failed with error: %v", err)
		c.AbortWithError(500, err)
	}
	encodedConfig, err := json.Marshal(config)
	if err != nil {
		fmt.Printf("Failed with error: %v", err)
		c.AbortWithError(500, err)
	}
	c.Writer.Write(encodedConfig)
}

func (ctrl *adminController) UpdateConfig(c *gin.Context) error {
	// TODO: add update config
	fmt.Println("Update Config")

	return nil
}

func (ctrl *adminController) GetCountryDetails(c *gin.Context) {
	config, err := ctrl.service.GetConfig()
	if err != nil {
		fmt.Printf("Failed with error: %v", err)
	}
	encodedConfig, err := json.Marshal(config)
	if err != nil {
		fmt.Printf("Failed with error: %v", err)
	}
	c.Writer.Write(encodedConfig)

}

func (ctrl *adminController) UpdateCountryDetails(c *gin.Context) error {
	// TODO: Take an object, update accordingly and return error
	fmt.Println("Update country details")

	return nil
}

func (ctrl *adminController) CreateCountryDetails(c *gin.Context) error {
	// TODO: Take an object, create and return error
	fmt.Println("Create details for new country")

	return nil
}

func (ctrl *adminController) DeleteCountry(c *gin.Context) error {
	// TODO: Take an object, create and return error
	fmt.Println("Delete by country details by id")

	return nil
}
