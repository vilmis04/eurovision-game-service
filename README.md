# Eurovision guessing game BE service built with Go

## Local setup

### Without database

To run locally just the service: `go run .`

To build an executable: `go build .`

### In docker with db and db adminer tool

To start virtual environment: `make start`

To stop virtual environment: `make stop`

To attach to go-service when container is running: `make attach`

## Configuration

The service must only be reachable through the auth proxy. Do not assign a public domain or a published port to it in Coolify.

| Env var | Required | Purpose |
|---|---|---|
| `INTERNAL_TOKEN` | yes | Shared secret. The proxy sends it in the `X-Internal-Token` header; requests without it are rejected with 401 and the `user` header is never trusted. |
| `INVITE_SECRET` | yes | HMAC key used to sign group invites (valid for 7 days). Changing it invalidates all outstanding invites. |
| `ADMIN_USERS` | no | Comma separated user ids allowed to create, update and delete countries. Empty means nobody. |

The service refuses to start without `INTERNAL_TOKEN` and `INVITE_SECRET`. `/api/health` is the only unauthenticated route.

## Checks

`make check` runs `go vet`, the tests and a scan that no SQL is built with `fmt.Sprintf`.

## Roadmap

- [x] proxy for permission checking
- [x] env var to define admin list
- [x] score calculation

## Deployment

- run command to deploy to image to dockerhub (replace x.x.x with version): `docker build -t vsud/ev-game:service-x.x.x . && docker push vsud/ev-game:service-x.x.x`
- ssh into the server, update docker compose with the new image version tag: `vi ~/PROJECTS/docker-compose.yaml`
- run `docker compose up -d` to start the service with the changes
