package country

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-service/internal/types"
	"github.com/vilmis04/eurovision-game-service/internal/utils"
)

type controller struct {
	service *Service
	router  *gin.RouterGroup
}

func NewController(app *gin.Engine) *controller {
	return &controller{
		service: NewService(),
		router:  app.Group("api/country"),
	}
}

func (ctrl *controller) Use() {
	ctrl.router.POST("/", func(c *gin.Context) {
		id, err := ctrl.service.CreateCountry(c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusCreated)
		c.Writer.Header().Set(types.HeaderContentType, types.HeaderApplicationJson)
		c.Writer.Write(*id)
	})

	ctrl.router.GET("/:year", func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		gameType := queryParams.Get("gameType")
		name := queryParams.Get("name")

		countries, err := ctrl.service.GetCountrySummary(c.Param("year"), gameType, name)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.Header().Set(types.HeaderContentType, types.HeaderApplicationJson)
		c.Writer.Write(*countries)
	})

	ctrl.router.PATCH(":year/:name", func(c *gin.Context) {
		params := map[string]string{
			"year": c.Param("year"),
			"name": c.Param("name"),
		}

		err := ctrl.service.UpdateCountry(params, c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.DELETE(":year/:name", func(c *gin.Context) {
		params := map[string]string{
			"year": c.Param("year"),
			"name": c.Param("name"),
		}

		err := ctrl.service.DeleteCountry(&params)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})
}
