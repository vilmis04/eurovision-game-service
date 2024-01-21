package admin

import "github.com/gin-gonic/gin"

func NewRouter(app *gin.Engine) {
	adminRoutes := app.Group("api/admin")
	adminController := NewController()

	// TODO: add auth middleware for admin endpoints
	adminRoutes.GET("/", func(c *gin.Context) {
		adminController.GetConfig(c)
	})

	adminRoutes.PATCH("/", func(c *gin.Context) {
		adminController.UpdateConfig(c)
	})
}
