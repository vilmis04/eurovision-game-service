package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/auth"
	"github.com/vilmis04/eurovision-game-service/internal/country"
	"github.com/vilmis04/eurovision-game-service/internal/db"
	"github.com/vilmis04/eurovision-game-service/internal/group"
	"github.com/vilmis04/eurovision-game-service/internal/health"
	"github.com/vilmis04/eurovision-game-service/internal/migrations"
	"github.com/vilmis04/eurovision-game-service/internal/score"
)

const shutdownTimeout = 15 * time.Second

func main() {
	// handled first: the probe needs only PORT, not the secrets and the database
	if len(os.Args) > 1 && os.Args[1] == "healthcheck" {
		os.Exit(runHealthcheck())
	}

	if err := run(); err != nil {
		log.Fatalf("[Server] %v", err)
	}
}

// requireEnv fails at startup instead of running without a secret.
func requireEnv(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", fmt.Errorf("%v must be set", name)
	}

	return value, nil
}

// portFromEnv returns the explicit listen port. There is no default.
func portFromEnv() (string, error) {
	port, err := requireEnv("PORT")
	if err != nil {
		return "", err
	}

	return validPort(port)
}

func validPort(port string) (string, error) {
	number, err := strconv.Atoi(port)
	if err != nil || number < 1 || number > 65535 {
		return "", errors.New("PORT must be a number between 1 and 65535")
	}

	return port, nil
}

func run() error {
	port, err := portFromEnv()
	if err != nil {
		return err
	}
	internalToken, err := requireEnv("INTERNAL_TOKEN")
	if err != nil {
		return err
	}
	if _, err := requireEnv("INVITE_SECRET"); err != nil {
		return err
	}
	dbConfig, err := db.ConfigFromEnv()
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	database, err := db.Open(ctx, dbConfig)
	if err != nil {
		return err
	}
	defer database.Close()

	if err := migrations.Run(ctx, database); err != nil {
		return err
	}

	adminService := admin.NewService(database)
	countryService := country.NewService(database, adminService)
	scoreService := score.NewService(database, adminService, countryService)
	groupService := group.NewService(database, scoreService)

	app := gin.Default()

	// health is registered before the proxy middleware so it stays unauthenticated
	app.GET("api/health", health.Handler(database))

	// every route registered below trusts the `user` header only when the
	// request carries the shared token that the auth proxy adds
	app.Use(auth.Proxy(internalToken))

	admin.NewController(app, adminService).Use()
	country.NewController(app, countryService).Use()
	group.NewController(app, groupService).Use()
	score.NewController(app, scoreService).Use()

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           app,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		log.Printf("[Server] listening on %v", server.Addr)
		serverErr <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
		log.Printf("[Server] shutting down")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}
