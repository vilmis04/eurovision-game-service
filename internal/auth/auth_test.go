package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func newRouter(token string, admins []string) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Proxy(token))
	router.GET("/who", func(c *gin.Context) { c.String(http.StatusOK, User(c)) })
	router.GET("/admin", Admin(admins), func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	return router
}

func do(router *gin.Engine, path string, headers map[string]string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	return rec
}

func TestProxy(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		headers map[string]string
		status  int
	}{
		{"missing token", "secret", map[string]string{"user": "bob"}, http.StatusUnauthorized},
		{"wrong token", "secret", map[string]string{HeaderInternalToken: "nope", "user": "bob"}, http.StatusUnauthorized},
		{"missing user", "secret", map[string]string{HeaderInternalToken: "secret"}, http.StatusUnauthorized},
		{"blank user", "secret", map[string]string{HeaderInternalToken: "secret", "user": "  "}, http.StatusUnauthorized},
		{"unconfigured token fails closed", "", map[string]string{HeaderInternalToken: "", "user": "bob"}, http.StatusUnauthorized},
		{"valid", "secret", map[string]string{HeaderInternalToken: "secret", "user": "bob"}, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := do(newRouter(tt.token, nil), "/who", tt.headers)
			if rec.Code != tt.status {
				t.Fatalf("status = %d, want %d", rec.Code, tt.status)
			}
			if tt.status == http.StatusOK && rec.Body.String() != "bob" {
				t.Fatalf("user = %q, want bob", rec.Body.String())
			}
		})
	}
}

func TestAdmin(t *testing.T) {
	router := newRouter("secret", []string{"alice"})

	rec := do(router, "/admin", map[string]string{HeaderInternalToken: "secret", "user": "bob"})
	if rec.Code != http.StatusForbidden {
		t.Fatalf("non admin status = %d, want 403", rec.Code)
	}

	rec = do(router, "/admin", map[string]string{HeaderInternalToken: "secret", "user": "alice"})
	if rec.Code != http.StatusOK {
		t.Fatalf("admin status = %d, want 200", rec.Code)
	}

	rec = do(newRouter("secret", nil), "/admin", map[string]string{HeaderInternalToken: "secret", "user": "alice"})
	if rec.Code != http.StatusForbidden {
		t.Fatalf("empty admin list status = %d, want 403", rec.Code)
	}
}

func TestParseList(t *testing.T) {
	got := ParseList(" a, b ,,c ")
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Fatalf("got %v", got)
	}
	if len(ParseList("")) != 0 {
		t.Fatal("empty input must give empty list")
	}
}
