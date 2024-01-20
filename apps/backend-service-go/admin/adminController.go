package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/utils"
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
	encodedConfig, err := ctrl.service.GetConfig()
	if err != nil {
		utils.HandleServerError(err, c)
		return
	}
	c.Writer.Write(*encodedConfig)
}

func (ctrl *adminController) UpdateConfig(c *gin.Context) {
	err := ctrl.service.UpdateConfig(c.Request)
	if err != nil {
		utils.HandleServerError(err, c)

		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func (ctrl *adminController) GetCountryDetails(c *gin.Context) {
	config, err := ctrl.service.GetConfig()
	if err != nil {
		utils.HandleServerError(err, c)
		return
	}
	encodedConfig, err := json.Marshal(config)
	if err != nil {
		utils.HandleServerError(err, c)
		return
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
