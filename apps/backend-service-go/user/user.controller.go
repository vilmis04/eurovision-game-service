package user

import (
	"github.com/gin-gonic/gin"
)

type userController struct {
	service *UserService
	router  *gin.RouterGroup
}

func NewController(app *gin.Engine) *userController {
	return &userController{
		service: NewService(),
		router:  app.Group("api/user"),
	}
}

func (ctrl *userController) Use() {
	ctrl.router.GET(":id", func(c *gin.Context) {
		ctrl.service.User(c.Param("id"))
	})

	ctrl.router.POST("/", func(c *gin.Context) {
		ctrl.service.NewUser(c.Request)
	})

	ctrl.router.PATCH(":id", func(c *gin.Context) {
		ctrl.service.UpdateUser(c.Param("id"), User{Roles: []string{"admin"}})
	})

	ctrl.router.DELETE(":id", func(c *gin.Context) {
		ctrl.service.DeleteUser(c.Param("id"))
	})
}
