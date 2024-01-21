package user

import "github.com/gin-gonic/gin"

func NewRouter(app *gin.Engine) {
	userRoutes := app.Group("api/user")
	userController := NewController()

	userRoutes.GET(":id", func(c *gin.Context) {
		userController.GetUser(c)
	})

	userRoutes.POST("/", func(c *gin.Context) {
		userController.NewUser(c)
	})

	userRoutes.PATCH(":id", func(c *gin.Context) {
		userController.UpdateUser(c)
	})

	userRoutes.DELETE(":id", func(c *gin.Context) {
		userController.DeleteUser(c)
	})
}
