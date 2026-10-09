// Package dbtest provides a sqlmock backed *sql.DB for repo tests.
package dbtest

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// NewMock returns a *sql.DB whose mock compares queries literally, so any
// value interpolated into the SQL text fails the expectation.
func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { database.Close() })

	return database, mock
}
