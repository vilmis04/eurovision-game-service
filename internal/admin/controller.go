package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-service/internal/utils"
)

type adminController struct {
	service *AdminService
	router  *gin.RouterGroup
}

func NewController(app *gin.Engine) *adminController {
	return &adminController{
		service: NewService(),
		router:  app.Group("api/admin"),
	}
}

func (ctrl *adminController) Use() {
	ctrl.router.GET("/", func(c *gin.Context) {
		encodedConfig, err := ctrl.service.GetConfig()
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}
		c.Writer.Write(*encodedConfig)
	})

	ctrl.router.PATCH("/", func(c *gin.Context) {
		err := ctrl.service.UpdateConfig(c.Request)
		if err != nil {
			utils.HandleServerError(err, c)

			return
		}
		c.Writer.WriteHeader(http.StatusOK)

	})
}
