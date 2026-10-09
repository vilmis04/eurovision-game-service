package dbtest

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/vilmis04/eurovision-game-service/internal/migrations"
)

// NewDatabase creates an empty, throwaway database on the server named by
// TEST_DATABASE_URL and returns a pool for it plus its URL (to open more pools).
// The test is skipped when TEST_DATABASE_URL is not set. The database is
// dropped when the test ends.
func NewDatabase(t *testing.T) (*sql.DB, string) {
	t.Helper()

	base := os.Getenv("TEST_DATABASE_URL")
	if base == "" {
		t.Skip("TEST_DATABASE_URL is not set")
	}

	baseURL, err := url.Parse(base)
	if err != nil {
		t.Fatalf("invalid TEST_DATABASE_URL: %v", err)
	}
	admin, err := sql.Open("postgres", base)
	if err != nil {
		t.Fatal(err)
	}

	name := fmt.Sprintf("evtest_%d", time.Now().UnixNano())
	if _, err := admin.Exec(`CREATE DATABASE ` + name); err != nil {
		admin.Close()
		t.Fatalf("create test database: %v", err)
	}

	testURL := *baseURL
	testURL.Path = "/" + name
	database, err := sql.Open("postgres", testURL.String())
	if err != nil {
		t.Fatal(err)
	}
	database.SetMaxOpenConns(10)

	t.Cleanup(func() {
		database.Close()
		_, _ = admin.Exec(`DROP DATABASE IF EXISTS ` + name + ` WITH (FORCE)`)
		admin.Close()
	})

	return database, testURL.String()
}

// NewMigrated is NewDatabase with every migration applied.
func NewMigrated(t *testing.T) *sql.DB {
	t.Helper()

	database, _ := NewDatabase(t)
	if err := migrations.Run(context.Background(), database); err != nil {
		t.Fatalf("migrate: %v", err)
	}

	return database
}
