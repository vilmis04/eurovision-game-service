package user

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *UserService
}

func NewController() *UserController {

	return &UserController{service: NewService()}
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	ctrl.service.User(c.Param("id"))
}

func (ctrl *UserController) NewUser(c *gin.Context) {
	ctrl.service.NewUser(c.Request)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	ctrl.service.UpdateUser(c.Param("id"), User{Roles: []string{"admin"}})
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	ctrl.service.DeleteUser(c.Param("id"))
}
