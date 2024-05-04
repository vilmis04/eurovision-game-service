package admin

import (
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
		router:  app.Group("api/admin"),
	}
}

func (ctrl *controller) Use() {
	ctrl.router.GET("/", func(c *gin.Context) {
		encodedConfig, err := ctrl.service.GetConfig()
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}
		c.Writer.Write(*encodedConfig)
	})

	// Disabled until admin portal is implemented

	// ctrl.router.PATCH("/", func(c *gin.Context) {
	// 	err := ctrl.service.UpdateConfig(c.Request)
	// 	if err != nil {
	// 		utils.HandleServerError(err, c)
	// 		return
	// 	}
	// 	c.Writer.WriteHeader(http.StatusOK)
	// })
}
