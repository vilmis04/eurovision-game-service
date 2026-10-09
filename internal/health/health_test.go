package health

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func serve(t *testing.T, pingErr error) *httptest.ResponseRecorder {
	t.Helper()
	gin.SetMode(gin.TestMode)

	database, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	if pingErr != nil {
		mock.ExpectPing().WillReturnError(pingErr)
	} else {
		mock.ExpectPing()
	}

	app := gin.New()
	app.GET("/api/health", Handler(database))
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/health", nil))

	return rec
}

func TestHealthOK(t *testing.T) {
	rec := serve(t, nil)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"OK"`) {
		t.Fatalf("status = %d, body = %q", rec.Code, rec.Body.String())
	}
}

func TestHealthUnavailableHidesCause(t *testing.T) {
	rec := serve(t, errors.New("pq: password authentication failed for user evgame"))
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "password") || strings.Contains(rec.Body.String(), "evgame") {
		t.Fatalf("cause leaked: %q", rec.Body.String())
	}
}
