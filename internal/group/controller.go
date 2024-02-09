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
	ctrl.router.GET(":owner", func(c *gin.Context) {
		// TODO: owner should come from cookie (in service maybe?), not from param
		groups, err := ctrl.service.GetGroups(c.Param("owner"), c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.Header().Set(types.HeaderContentType, types.HeaderApplicationJson)
		c.Writer.Write(*groups)
	})

	ctrl.router.POST(":owner", func(c *gin.Context) {
		// TODO: owner should come from cookie (in service maybe?), not from param
		id, err := ctrl.service.CreateGroup(c.Param("owner"), c.Request)
		if err != nil {
			utils.HandleServerError(err, c)
			return
		}

		c.Writer.WriteHeader(http.StatusCreated)
		c.Writer.Write(*id)
	})
}
