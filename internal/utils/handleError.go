package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleServerError(err error, c *gin.Context) {
	c.AbortWithError(http.StatusInternalServerError, err)
	log.Println(err)
	c.Writer.Write([]byte(err.Error()))
}

func HandleClientError(err error, c *gin.Context) {
	c.AbortWithError(http.StatusBadRequest, err)
	log.Println(err)
	c.Writer.Write([]byte(err.Error()))
}
