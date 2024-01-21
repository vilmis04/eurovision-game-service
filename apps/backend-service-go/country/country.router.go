package country

import "github.com/gin-gonic/gin"

func NewRouter(app *gin.Engine) {
	countryRoutes := app.Group("api/country")
	countryController := NewController()

	countryRoutes.GET("/", func(c *gin.Context) {
		countryController.GetCountry(c)
	})

	countryRoutes.GET(":id", func(c *gin.Context) {
		countryController.GetCountry(c)
	})

	countryRoutes.PUT(":id", func(c *gin.Context) {
		countryController.UpdateCountry(c)
	})

	countryRoutes.POST(":id", func(c *gin.Context) {
		countryController.CreateCountry(c)
	})

	countryRoutes.DELETE(":id", func(c *gin.Context) {
		countryController.DeleteCountry(c)
	})
}
