// Package auth enforces the trust boundary between the auth proxy and this service.
//
// The service is only reachable through the proxy. The proxy authenticates the
// caller and forwards the identity in the `user` header together with a shared
// secret in the `X-Internal-Token` header. Without the secret the `user`
// header is not trusted and the request is rejected.
package auth

import (
	"crypto/subtle"
	"os"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-service/internal/utils"
)

const (
	HeaderInternalToken = "X-Internal-Token"
	HeaderUser          = "user"

	contextUserKey = "auth.user"
)

// Proxy rejects requests that do not carry the shared INTERNAL_TOKEN or that
// have no user, and stores the verified user in the request context.
// An empty token rejects everything so a misconfigured service fails closed.
func Proxy(token string) gin.HandlerFunc {
	expected := []byte(token)

	return func(c *gin.Context) {
		provided := []byte(c.GetHeader(HeaderInternalToken))
		if len(expected) == 0 || subtle.ConstantTimeCompare(provided, expected) != 1 {
			utils.HandleError(utils.Unauthorized("unauthorized"), c)
			return
		}

		user := strings.TrimSpace(c.GetHeader(HeaderUser))
		if user == "" {
			utils.HandleError(utils.Unauthorized("unauthorized"), c)
			return
		}

		c.Set(contextUserKey, user)
		c.Next()
	}
}

// User returns the user verified by the Proxy middleware, or "" when absent.
func User(c *gin.Context) string {
	return c.GetString(contextUserKey)
}

// Admin allows only users listed in the admins list. Must run after Proxy.
func Admin(admins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := User(c)
		if user == "" || !slices.Contains(admins, user) {
			utils.HandleError(utils.Forbidden("forbidden"), c)
			return
		}

		c.Next()
	}
}

// ParseList splits a comma separated list, trimming blanks and empty entries.
func ParseList(value string) []string {
	list := []string{}
	for _, item := range strings.Split(value, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			list = append(list, item)
		}
	}

	return list
}

// AdminsFromEnv reads the ADMIN_USERS env var (comma separated user ids).
func AdminsFromEnv() []string {
	return ParseList(os.Getenv("ADMIN_USERS"))
}
