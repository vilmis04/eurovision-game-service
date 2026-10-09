// Package health serves the unauthenticated liveness and readiness endpoint.
package health

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const pingTimeout = 2 * time.Second

// Pinger is satisfied by *sql.DB.
type Pinger interface {
	PingContext(ctx context.Context) error
}

// Handler answers 200 when the database responds and 503 otherwise.
// The failure reason is never sent to the client.
func Handler(database Pinger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), pingTimeout)
		defer cancel()

		if err := database.PingContext(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"health": "unavailable"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"health": "OK"})
	}
}
