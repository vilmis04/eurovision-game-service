package admin

import (
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
