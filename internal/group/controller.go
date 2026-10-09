package group

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-service/internal/auth"
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
		groups, err := ctrl.service.GetGroups(auth.User(c), c.Request)
		if err != nil {
			utils.HandleError(err, c)
			return
		}

		c.Writer.Header().Set(types.HeaderContentType, types.HeaderApplicationJson)
		c.Writer.Write(*groups)
	})

	ctrl.router.POST("/", func(c *gin.Context) {
		id, err := ctrl.service.CreateGroup(auth.User(c), c.Request)
		if err != nil {
			utils.HandleError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusCreated)
		c.Writer.Write(*id)
	})

	ctrl.router.PATCH(":id", func(c *gin.Context) {
		err := ctrl.service.UpdateMembers(auth.User(c), c.Param("id"), c.Request)
		if err != nil {
			utils.HandleError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.DELETE(":id", func(c *gin.Context) {
		err := ctrl.service.DeleteGroup(auth.User(c), c.Param("id"))
		if err != nil {
			utils.HandleError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.POST(":id/generate-invite", func(c *gin.Context) {
		inviteCode, err := ctrl.service.GenerateInvite(c.Param("id"), auth.User(c))
		if err != nil {
			utils.HandleError(err, c)
			return
		}

		c.Writer.Header().Set(types.HeaderContentType, "text/plain")
		c.String(http.StatusCreated, inviteCode)
	})

	ctrl.router.POST("join", func(c *gin.Context) {
		err := ctrl.service.JoinGroup(auth.User(c), c.Request)
		if err != nil {
			utils.HandleError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	})

	ctrl.router.GET("leaderboard", func(c *gin.Context) {
		leaderboard, err := ctrl.service.GetLeaderboard(auth.User(c), c.Query("id"))
		if err != nil {
			utils.HandleError(err, c)
			return
		}

		c.Writer.Header().Set(types.HeaderContentType, types.HeaderApplicationJson)
		c.Writer.Write(*leaderboard)
	})
}
