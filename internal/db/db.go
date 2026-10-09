// Package db builds the single shared connection pool used by every repo.
package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
	"slices"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	defaultSSLMode         = "require" // lib/pq's own default, kept explicit
	defaultMaxOpenConns    = 10
	defaultMaxIdleConns    = 10
	defaultConnMaxLifetime = 30 * time.Minute
	defaultConnMaxIdleTime = 5 * time.Minute
	defaultConnectTimeout  = 30 * time.Second

	// migrations hold one connection for the advisory lock and use another to run
	minOpenConns = 2
)

var sslModes = []string{"disable", "allow", "prefer", "require", "verify-ca", "verify-full"}

// Config holds everything needed to open the pool. It deliberately has no
// String method and no error ever contains the URL, because it carries the password.
type Config struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	ConnectTimeout  time.Duration
}

// ConfigFromEnv reads DATABASE_URL, DB_SSLMODE and the DB_* pool settings.
// sslmode in the URL wins over DB_SSLMODE, which defaults to "require".
func ConfigFromEnv() (Config, error) {
	return configFrom(os.Getenv)
}

func configFrom(getenv func(string) string) (Config, error) {
	raw := getenv("DATABASE_URL")
	if raw == "" {
		return Config{}, errors.New("DATABASE_URL must be set")
	}

	parsed, err := url.Parse(raw)
	if err != nil || (parsed.Scheme != "postgres" && parsed.Scheme != "postgresql") || parsed.Host == "" {
		return Config{}, errors.New("DATABASE_URL must be a postgres:// URL with a host")
	}

	query := parsed.Query()
	if query.Get("sslmode") == "" {
		mode := getenv("DB_SSLMODE")
		if mode == "" {
			mode = defaultSSLMode
		}
		query.Set("sslmode", mode)
	}
	if !slices.Contains(sslModes, query.Get("sslmode")) {
		return Config{}, fmt.Errorf("sslmode must be one of %v", sslModes)
	}
	parsed.RawQuery = query.Encode()

	cfg := Config{
		DSN:             parsed.String(),
		MaxOpenConns:    defaultMaxOpenConns,
		MaxIdleConns:    defaultMaxIdleConns,
		ConnMaxLifetime: defaultConnMaxLifetime,
		ConnMaxIdleTime: defaultConnMaxIdleTime,
		ConnectTimeout:  defaultConnectTimeout,
	}
	if err := intFrom(getenv, "DB_MAX_OPEN_CONNS", &cfg.MaxOpenConns); err != nil {
		return Config{}, err
	}
	if err := intFrom(getenv, "DB_MAX_IDLE_CONNS", &cfg.MaxIdleConns); err != nil {
		return Config{}, err
	}
	if err := durationFrom(getenv, "DB_CONN_MAX_LIFETIME", &cfg.ConnMaxLifetime); err != nil {
		return Config{}, err
	}
	if err := durationFrom(getenv, "DB_CONN_MAX_IDLE_TIME", &cfg.ConnMaxIdleTime); err != nil {
		return Config{}, err
	}
	if err := durationFrom(getenv, "DB_CONNECT_TIMEOUT", &cfg.ConnectTimeout); err != nil {
		return Config{}, err
	}

	if cfg.MaxOpenConns < minOpenConns {
		return Config{}, fmt.Errorf("DB_MAX_OPEN_CONNS must be at least %d", minOpenConns)
	}
	if cfg.MaxIdleConns > cfg.MaxOpenConns {
		cfg.MaxIdleConns = cfg.MaxOpenConns
	}

	return cfg, nil
}

func intFrom(getenv func(string) string, name string, target *int) error {
	value := getenv(name)
	if value == "" {
		return nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed < 0 {
		return fmt.Errorf("%s must be a non-negative integer", name)
	}
	*target = parsed

	return nil
}

func durationFrom(getenv func(string) string, name string, target *time.Duration) error {
	value := getenv(name)
	if value == "" {
		return nil
	}
	parsed, err := time.ParseDuration(value)
	if err != nil || parsed < 0 {
		return fmt.Errorf("%s must be a duration such as 30m", name)
	}
	*target = parsed

	return nil
}

// Open creates the pool and waits until the database answers, retrying with
// backoff for up to ConnectTimeout because Postgres may still be starting.
func Open(ctx context.Context, cfg Config) (*sql.DB, error) {
	database, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, errors.New("invalid database configuration")
	}

	database.SetMaxOpenConns(cfg.MaxOpenConns)
	database.SetMaxIdleConns(cfg.MaxIdleConns)
	database.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	database.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	if err := waitForDatabase(ctx, database, cfg.ConnectTimeout); err != nil {
		database.Close()
		return nil, err
	}

	return database, nil
}

func waitForDatabase(ctx context.Context, database *sql.DB, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	delay := 500 * time.Millisecond
	for attempt := 1; ; attempt++ {
		pingCtx, pingCancel := context.WithTimeout(ctx, 3*time.Second)
		err := database.PingContext(pingCtx)
		pingCancel()
		if err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("database not reachable after %d attempts", attempt)
		case <-time.After(delay):
		}
		if delay < 3*time.Second {
			delay *= 2
		}
	}
}
