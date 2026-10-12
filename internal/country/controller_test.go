package country

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/auth"
)

// The service has no database behind it here: every request below must be
// rejected before any query would run.
func newTestRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	t.Setenv("ADMIN_USERS", "admin")

	app := gin.New()
	app.Use(auth.Proxy("secret"))
	NewController(app, NewService(nil, admin.NewService(nil))).Use()

	return app
}

func request(app *gin.Engine, method string, path string, user string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(auth.HeaderInternalToken, "secret")
	req.Header.Set(auth.HeaderUser, user)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	return rec
}

func TestWriteEndpointsRequireAdmin(t *testing.T) {
	app := newTestRouter(t)
	createBody := `{"name":"Serbia","code":"rs","gameType":"semi1","artist":"a","song":"s","orderSemi":1}`

	tests := []struct{ method, path, body string }{
		{http.MethodPost, "/api/country/", createBody},
		{http.MethodPatch, "/api/country/2026/Serbia", `{"score":5}`},
		{http.MethodDelete, "/api/country/2026/Serbia", ""},
	}
	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			rec := request(app, tt.method, tt.path, "mallory", tt.body)
			if rec.Code != http.StatusForbidden {
				t.Fatalf("status = %d, want 403", rec.Code)
			}
		})
	}
}

func TestWriteEndpointsRejectMalformedRequests(t *testing.T) {
	app := newTestRouter(t)

	tests := []struct{ name, method, path, body string }{
		{"create with broken json", http.MethodPost, "/api/country/", `{`},
		{"create with missing fields", http.MethodPost, "/api/country/", `{"name":"Serbia"}`},
		{"create with invalid game type", http.MethodPost, "/api/country/", `{"name":"S","code":"rs","gameType":"x","artist":"a","song":"s","orderSemi":1}`},
		{"update with non numeric year", http.MethodPatch, "/api/country/1%20OR%201=1/Serbia", `{"score":5}`},
		{"update with empty body", http.MethodPatch, "/api/country/2026/Serbia", `{}`},
		{"update with broken json", http.MethodPatch, "/api/country/2026/Serbia", `{`},
		{"delete with non numeric year", http.MethodDelete, "/api/country/abc/Serbia", ""},
		{"read with non numeric year", http.MethodGet, "/api/country/abc", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := request(app, tt.method, tt.path, "admin", tt.body)
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400 (body %q)", rec.Code, rec.Body.String())
			}
		})
	}
}

func TestRequestsWithoutProxyTokenAreRejected(t *testing.T) {
	app := newTestRouter(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/country/2026/Serbia", nil)
	req.Header.Set(auth.HeaderUser, "admin") // spoofed identity, no token
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", rec.Code)
	}
}
