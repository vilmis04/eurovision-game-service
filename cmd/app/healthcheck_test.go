package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func portOf(t *testing.T, server *httptest.Server) string {
	t.Helper()
	_, port, err := net.SplitHostPort(strings.TrimPrefix(server.URL, "http://"))
	if err != nil {
		t.Fatal(err)
	}

	return port
}

func TestHealthcheckPassesOnOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/health" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	if err := healthcheck(server.Client(), portOf(t, server)); err != nil {
		t.Fatal(err)
	}
}

func TestHealthcheckFailsOnUnavailable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	err := healthcheck(server.Client(), portOf(t, server))
	if err == nil || !strings.Contains(err.Error(), "503") {
		t.Fatalf("err = %v, want a 503 failure", err)
	}
}

func TestHealthcheckFailsWhenNothingListens(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	port := strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	listener.Close()

	if err := healthcheck(http.DefaultClient, port); err == nil {
		t.Fatal("want an error when the service is not running")
	}
}

func TestHealthcheckRejectsBadPort(t *testing.T) {
	for _, port := range []string{"", "abc", "0", "70000", "80/../x"} {
		if err := healthcheck(http.DefaultClient, port); err == nil || !strings.Contains(err.Error(), "PORT") {
			t.Errorf("port %q: err = %v, want a PORT error", port, err)
		}
	}
}
