package group

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
		router:  app.Group("api/group"),
	}
}

func (ctrl *controller) Use() {
	ctrl.router.GET("/", func(c *gin.Context) {
		groups, err := ctrl.service.GetGroups(c.GetHeader("user"), c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.Header().Set(types.HeaderContentType, types.HeaderApplicationJson)
		c.Writer.Write(*groups)
	})

	ctrl.router.POST("/", func(c *gin.Context) {
		id, err := ctrl.service.CreateGroup(c.GetHeader("user"), c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusCreated)
		c.Writer.Write(*id)
	})

	ctrl.router.PATCH("/", func(c *gin.Context) {
		err := ctrl.service.UpdateMembers(c.GetHeader("user"), c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.DELETE(":name", func(c *gin.Context) {
		err := ctrl.service.DeleteGroup(c.GetHeader("user"), c.Param("name"))
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.POST(":name/generate-invite", func(c *gin.Context) {
		inviteCode, err := ctrl.service.GenerateInvite(c.Param("name"), c.GetHeader("user"))
		if err != nil {
			utils.HandleServerError(err, c)
		}

		c.Writer.Header().Set("Content-Type", "text/plain")
		c.String(http.StatusCreated, inviteCode)
	})

	ctrl.router.POST("join", func(c *gin.Context) {
		err := ctrl.service.JoinGroup(c.GetHeader("user"), c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
		}

		c.Writer.WriteHeader(http.StatusOK)
	})
}
