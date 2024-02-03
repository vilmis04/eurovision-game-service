package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleServerError(err error, c *gin.Context) {
	c.AbortWithError(500, err)
	c.Writer.Write([]byte(err.Error()))
}
