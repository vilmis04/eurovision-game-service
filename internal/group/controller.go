package group

import (
	"fmt"
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

	ctrl.router.PATCH(":owner/:name", func(c *gin.Context) {
		err := ctrl.service.UpdateMembers(c.Param("owner"), c.Param("name"), c.Request)
		if err != nil {
			if err.Error() == fmt.Sprint(http.StatusBadRequest) {
				utils.HandleClientError(err, c)
				return
			}
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.DELETE(":id", func(c *gin.Context) {
		err := ctrl.service.DeleteGroup(c.GetHeader("user"), c.Param("id"))
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.POST(":id/generate-invite", func(c *gin.Context) {
		inviteCode, err := ctrl.service.GenerateInvite(c.Param("id"), c.GetHeader("user"))
		if err != nil {
			utils.HandleServerError(err, c)
		}

		c.Writer.Header().Set(types.HeaderContentType, "text/plain")
		c.String(http.StatusCreated, inviteCode)
	})

	ctrl.router.POST("join", func(c *gin.Context) {
		err := ctrl.service.JoinGroup(c.GetHeader("user"), c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.GET("leaderboard", func(c *gin.Context) {
		leaderboard, err := ctrl.service.GetLeaderboard(c.GetHeader("user"), c.Query("id"))
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.Header().Set(types.HeaderContentType, types.HeaderApplicationJson)
		c.Writer.Write(*leaderboard)
	})
}
