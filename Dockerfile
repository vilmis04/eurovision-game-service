# syntax=docker/dockerfile:1.7

# Keep GO_VERSION in step with the `go` line in go.mod (CI checks this).
ARG GO_VERSION=1.24

FROM golang:${GO_VERSION}-alpine AS build
WORKDIR /src

# Dependencies first: this layer is reused until go.mod or go.sum change.
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Only what the build needs, so editing docs or workflows does not invalidate it.
COPY cmd ./cmd
COPY internal ./internal
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/app


# Static binary, CA certificates and tzdata, no shell, runs as non-root.
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/app /app

# Default so the image runs and passes its health check without configuration.
# Coolify's exposed port must equal PORT.
ENV PORT=8080
EXPOSE 8080

USER nonroot:nonroot

# No curl in the image: the binary probes /api/health itself.
HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 CMD ["/app", "healthcheck"]

ENTRYPOINT ["/app"]
