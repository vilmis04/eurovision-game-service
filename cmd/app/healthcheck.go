package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

const healthcheckTimeout = 3 * time.Second

// runHealthcheck is the `app healthcheck` subcommand used by the Docker HEALTHCHECK.
// The final image has no shell, curl or wget, so the binary probes itself.
// It exits 0 when /api/health answers 200 and 1 otherwise.
func runHealthcheck() int {
	if err := healthcheck(http.DefaultClient, os.Getenv("PORT")); err != nil {
		fmt.Fprintf(os.Stderr, "healthcheck failed: %v\n", err)
		return 1
	}

	return 0
}

func healthcheck(client *http.Client, port string) error {
	if _, err := validPort(port); err != nil {
		return err
	}

	return probe(client, "http://127.0.0.1:"+port+"/api/health")
}

func probe(client *http.Client, url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), healthcheckTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("service is not reachable")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	return nil
}
