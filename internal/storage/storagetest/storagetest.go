// Package storagetest provides a sqlmock backed storage for repo tests.
package storagetest

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vilmis04/eurovision-game-service/internal/storage"
)

// New returns a Storage wired to a sqlmock that compares queries literally,
// so any value interpolated into the SQL text fails the expectation.
func New(t *testing.T, table string) (*storage.Storage, sqlmock.Sqlmock) {
	t.Helper()

	dsn := t.Name()
	db, mock, err := sqlmock.NewWithDSN(dsn, sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })

	return &storage.Storage{ConnString: dsn, Table: table, Driver: "sqlmock"}, mock
}
