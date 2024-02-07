package group

import (
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
	ctrl.router.GET("/:owner", func(c *gin.Context) {
		groups, err := ctrl.service.GetGroups(c.Param("owner"), c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.Header().Set(types.HeaderContentType, types.HeaderApplicationJson)
		c.Writer.Write(*groups)
	})
}
