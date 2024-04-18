package score

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-service/internal/utils"
)

type controller struct {
	service *Service
	router  *gin.RouterGroup
}

func NewController(app *gin.Engine) *controller {
	return &controller{
		service: NewService(),
		router:  app.Group("api/score"),
	}
}

func (ctrl *controller) Use() {
	ctrl.router.PUT(":country", func(c *gin.Context) {
		country := c.Param("country")
		user := c.GetHeader("user")

		err := ctrl.service.UpdateScore(country, user, c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

}
