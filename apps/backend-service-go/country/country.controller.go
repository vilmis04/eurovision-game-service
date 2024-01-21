package country

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/utils"
)

type countryController struct {
	service *countryService
}

func NewController() *countryController {
	return &countryController{
		service: NewService(),
	}
}

func (ctrl *countryController) GetCountry(c *gin.Context) {
	config, err := ctrl.service.GetCountry()
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

func (ctrl *countryController) UpdateCountry(c *gin.Context) error {
	// TODO: Take an object, update accordingly and return error
	fmt.Println("Update country details")

	return nil
}

func (ctrl *countryController) CreateCountry(c *gin.Context) error {
	// TODO: Take an object, create and return error
	fmt.Println("Create details for new country")

	return nil
}

func (ctrl *countryController) DeleteCountry(c *gin.Context) error {
	// TODO: Take an object, create and return error
	fmt.Println("Delete by country details by id")

	return nil
}
