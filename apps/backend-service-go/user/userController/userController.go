package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/user/userModel"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/user/userService"
)

type UserController struct {
	service *userService.UserService
}

func New() *UserController {

	return &UserController{service: userService.New()}
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	ctrl.service.User(c.Param("id"))
}

func (ctrl *UserController) NewUser(c *gin.Context) {

	ctrl.service.NewUser(c.Request)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	ctrl.service.UpdateUser(c.Param("id"), userModel.User{Roles: []string{"admin"}})
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	ctrl.service.DeleteUser(c.Param("id"))
}
