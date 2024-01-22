package country

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/utils"
)

type countryController struct {
	service *CountryService
	router  *gin.RouterGroup
}

func NewController(app *gin.Engine) *countryController {
	return &countryController{
		service: NewService(),
		router:  app.Group("api/country"),
	}
}

func (ctrl *countryController) Use() {
	ctrl.router.POST("/", func(c *gin.Context) {
		id, err := ctrl.service.CreateCountry(c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusCreated)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(*id)
	})

	ctrl.router.GET("/", func(c *gin.Context) {
		config, err := ctrl.service.GetAllCountries()
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
	})

	ctrl.router.GET(":id", func(c *gin.Context) {
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
	})

	ctrl.router.PUT(":id", func(c *gin.Context) {
		err := ctrl.service.UpdateCountry()
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.DELETE(":id", func(c *gin.Context) {
		err := ctrl.service.DeleteCountry()
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})
}
