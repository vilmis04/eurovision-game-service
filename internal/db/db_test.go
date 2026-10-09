package db

import (
	"net/url"
	"strings"
	"testing"
	"time"
)

func env(values map[string]string) func(string) string {
	return func(name string) string { return values[name] }
}

func TestConfigFromEnv(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		wantErr string
		check   func(t *testing.T, cfg Config)
	}{
		{name: "missing url", env: map[string]string{}, wantErr: "DATABASE_URL must be set"},
		{name: "wrong scheme", env: map[string]string{"DATABASE_URL": "mysql://u:p@h/db"}, wantErr: "postgres:// URL"},
		{name: "no host", env: map[string]string{"DATABASE_URL": "postgres:///db"}, wantErr: "postgres:// URL"},
		{name: "invalid sslmode", env: map[string]string{"DATABASE_URL": "postgres://u:p@h/db", "DB_SSLMODE": "yes"}, wantErr: "sslmode must be one of"},
		{name: "too few connections", env: map[string]string{"DATABASE_URL": "postgres://u:p@h/db", "DB_MAX_OPEN_CONNS": "1"}, wantErr: "at least 2"},
		{name: "bad duration", env: map[string]string{"DATABASE_URL": "postgres://u:p@h/db", "DB_CONN_MAX_LIFETIME": "soon"}, wantErr: "DB_CONN_MAX_LIFETIME"},
		{name: "bad integer", env: map[string]string{"DATABASE_URL": "postgres://u:p@h/db", "DB_MAX_OPEN_CONNS": "many"}, wantErr: "DB_MAX_OPEN_CONNS"},
		{
			name: "defaults to sslmode require",
			env:  map[string]string{"DATABASE_URL": "postgres://u:p@h:5432/db"},
			check: func(t *testing.T, cfg Config) {
				if got := sslmode(t, cfg); got != "require" {
					t.Fatalf("sslmode = %q, want require", got)
				}
				if cfg.MaxOpenConns != 10 || cfg.ConnMaxLifetime != 30*time.Minute {
					t.Fatalf("unexpected defaults %+v", cfg)
				}
			},
		},
		{
			name: "DB_SSLMODE applies when url has none",
			env:  map[string]string{"DATABASE_URL": "postgres://u:p@h/db", "DB_SSLMODE": "disable"},
			check: func(t *testing.T, cfg Config) {
				if got := sslmode(t, cfg); got != "disable" {
					t.Fatalf("sslmode = %q, want disable", got)
				}
			},
		},
		{
			name: "sslmode in url wins",
			env:  map[string]string{"DATABASE_URL": "postgres://u:p@h/db?sslmode=verify-full", "DB_SSLMODE": "disable"},
			check: func(t *testing.T, cfg Config) {
				if got := sslmode(t, cfg); got != "verify-full" {
					t.Fatalf("sslmode = %q, want verify-full", got)
				}
			},
		},
		{
			name: "pool settings and idle clamp",
			env: map[string]string{
				"DATABASE_URL": "postgres://u:p@h/db", "DB_MAX_OPEN_CONNS": "4", "DB_MAX_IDLE_CONNS": "9",
				"DB_CONN_MAX_LIFETIME": "10m", "DB_CONN_MAX_IDLE_TIME": "1m",
			},
			check: func(t *testing.T, cfg Config) {
				if cfg.MaxOpenConns != 4 || cfg.MaxIdleConns != 4 || cfg.ConnMaxLifetime != 10*time.Minute || cfg.ConnMaxIdleTime != time.Minute {
					t.Fatalf("unexpected config %+v", cfg)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := configFrom(env(tt.env))
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("err = %v, want containing %q", err, tt.wantErr)
				}
				if strings.Contains(err.Error(), "u:p") {
					t.Fatalf("error leaks credentials: %v", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			tt.check(t, cfg)
		})
	}
}

func sslmode(t *testing.T, cfg Config) string {
	t.Helper()
	parsed, err := url.Parse(cfg.DSN)
	if err != nil {
		t.Fatal(err)
	}

	return parsed.Query().Get("sslmode")
}
