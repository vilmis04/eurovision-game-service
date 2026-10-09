// Package migrations embeds the versioned SQL migrations and applies them.
package migrations

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"

	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/lock"
)

//go:embed *.sql
var files embed.FS

// FS returns the embedded migration files.
func FS() fs.FS {
	return files
}

// NewProvider builds a goose provider over the shared pool. When locked, a
// Postgres session advisory lock is held for the duration of a run, so two
// instances starting at the same time (a rolling deploy) cannot apply the same
// migration twice: the second waits, then finds nothing left to do.
// The lock keeps one pooled connection busy, so the pool needs at least 2.
func NewProvider(database *sql.DB, locked bool) (*goose.Provider, error) {
	options := []goose.ProviderOption{}
	if locked {
		locker, err := lock.NewPostgresSessionLocker()
		if err != nil {
			return nil, fmt.Errorf("create migration lock: %w", err)
		}
		options = append(options, goose.WithSessionLocker(locker))
	}

	return goose.NewProvider(goose.DialectPostgres, database, files, options...)
}

// Run applies all pending migrations and logs what it did.
func Run(ctx context.Context, database *sql.DB) error {
	provider, err := NewProvider(database, true)
	if err != nil {
		return fmt.Errorf("prepare migrations: %w", err)
	}

	results, err := provider.Up(ctx)
	if err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}
	for _, result := range results {
		log.Printf("[migrations] applied %s in %s", result.Source.Path, result.Duration)
	}
	if len(results) == 0 {
		log.Printf("[migrations] database is up to date")
	}

	return nil
}
