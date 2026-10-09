package utils

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPError is an error whose message is safe to show to API clients.
// Anything that is not an HTTPError is treated as internal and never exposed.
type HTTPError struct {
	Status  int
	Message string
}

func (e *HTTPError) Error() string {
	return e.Message
}

func BadRequest(message string) error {
	return &HTTPError{Status: http.StatusBadRequest, Message: message}
}

func Unauthorized(message string) error {
	return &HTTPError{Status: http.StatusUnauthorized, Message: message}
}

func Forbidden(message string) error {
	return &HTTPError{Status: http.StatusForbidden, Message: message}
}

func NotFound(message string) error {
	return &HTTPError{Status: http.StatusNotFound, Message: message}
}

func Conflict(message string) error {
	return &HTTPError{Status: http.StatusConflict, Message: message}
}

// HandleError logs the error and answers with a client-safe body.
// Only HTTPError messages reach the client, everything else becomes a generic 500.
func HandleError(err error, c *gin.Context) {
	log.Println(err)

	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		c.AbortWithStatusJSON(httpErr.Status, gin.H{"error": httpErr.Message})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
