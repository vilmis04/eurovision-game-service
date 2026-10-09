# Eurovision guessing game BE service built with Go

## Local setup

Prerequisites: the Go version in `go.mod` and Docker with Compose (for Postgres).

1. Copy the example environment file and adjust it if you like: `cp .env.example .env`.
   `.env` is ignored by git. The service itself does not read `.env` files; `make run` and
   `docker compose` load it for you.
2. Start only the database: `docker compose up -d db`. It listens on `localhost:5432`
   (Adminer is available at `http://localhost:5555` with `docker compose up -d admin`).
3. Run the service: `make run` (loads `.env`, then `go run ./cmd/app`). Without make:
   `set -a; . ./.env; set +a; go run ./cmd/app`.
4. On the first start the schema and the country data are created by the migrations, so an
   empty database is ready without any extra step.
5. Try it. Only `/api/health` is open; everything else needs what the auth proxy normally adds:

```sh
curl localhost:4200/api/health
curl -H 'X-Internal-Token: change-me-local-token' -H 'user: me' localhost:4200/api/score/
```

Tests: `go test ./...` needs no database. The integration tests (migrations and score rows)
run against a real Postgres when `TEST_DATABASE_URL` is set to a superuser URL, for example
`TEST_DATABASE_URL='postgres://evgame:evgame@localhost:5432/postgres?sslmode=disable' go test ./...`.
They create and drop their own throwaway databases and are skipped when the variable is unset.

Reset the local database (drops the data volume, the next start rebuilds the schema):
`docker compose down -v`.

Other make targets: `make start` / `make stop` / `make attach` run the whole stack in Docker.

### Troubleshooting

- `PORT must be set` (or `INTERNAL_TOKEN`, `INVITE_SECRET`, `DATABASE_URL`): the variable is missing from `.env` or
  the shell. The service refuses to start without them.
- `pq: SSL is not enabled on the server`: set `DB_SSLMODE=disable` (the compose database has no TLS).
- `database not reachable after N attempts`: Postgres is not running or `DATABASE_URL` is wrong. The service
  retries for `DB_CONNECT_TIMEOUT` (30 s by default) so a database that is still starting is fine.
- `address already in use`: change `PORT`.

## Configuration

The service must only be reachable through the auth proxy. Do not assign a public domain or a published port to it in Coolify.

| Env var | Required | Purpose |
|---|---|---|
| `PORT` | yes | TCP port to listen on. There is no default. |
| `INTERNAL_TOKEN` | yes | Shared secret. The proxy sends it in the `X-Internal-Token` header; requests without it are rejected with 401 and the `user` header is never trusted. |
| `INVITE_SECRET` | yes | HMAC key used to sign group invites (valid for 7 days). Changing it invalidates all outstanding invites. |
| `ADMIN_USERS` | no | Comma separated user ids allowed to create, update and delete countries. Empty means nobody. |
| `DATABASE_URL` | yes | Postgres URL, `postgres://user:password@host:5432/dbname`. |
| `DB_SSLMODE` | no | `sslmode` used when the URL has none. Default `require`; set `disable` on an internal Coolify network without TLS. |
| `DB_MAX_OPEN_CONNS` | no | Pool size, default 10, minimum 2 (migrations hold one connection for their lock). |
| `DB_MAX_IDLE_CONNS` | no | Idle connections kept, default 10. |
| `DB_CONN_MAX_LIFETIME` | no | Maximum age of a connection, default `30m`. |
| `DB_CONN_MAX_IDLE_TIME` | no | Maximum idle time of a connection, default `5m`. |
| `DB_CONNECT_TIMEOUT` | no | How long startup waits for the database, default `30s`. |

The service refuses to start when a required variable is missing or invalid. `/api/health` is the only
unauthenticated route; it answers 200 when the database responds and 503 otherwise.

## Database migrations

Schema and seed data live in versioned SQL files in `internal/migrations`, embedded into the binary
([goose](https://github.com/pressly/goose)). They are applied automatically at startup, before the HTTP server
starts, while holding a Postgres advisory lock. When two instances start together (rolling deploy) the second
one waits, finds nothing left to apply and carries on. A failing migration stops the service from starting, so
the deploy is marked failed instead of running on a half-migrated database. Each migration runs in a transaction.

Files, in order:

| File | What it does |
|---|---|
| `0001_baseline.sql` | The tables, all `IF NOT EXISTS`, so a database created earlier adopts it unchanged. |
| `0002_integrity.sql` | Collapses old duplicates, adds unique keys (`score(user, country, year)`, `country(year, name)`, `group(owner, name)`), indexes, and makes sure `group.members` is a real `text[]`. |
| `0003_seed_admin_config.sql` | The single configuration row (year 2026, `semi1`, voting off). |
| `0004`-`0006_seed_countries_<year>.sql` | The countries of each contest year. Existing rows are never overwritten. |

### Add a migration

1. `make migrate-new name=add_something` creates the next numbered file (`00007_add_something.sql`).
2. Write the `-- +goose Up` and `-- +goose Down` sections. For a `DO $$ ... $$` block wrap it in
   `-- +goose StatementBegin` / `-- +goose StatementEnd`.
3. Run `make migrate-up` against your local database, then `go test ./...` with `TEST_DATABASE_URL` set.
4. Commit it.

A new contest year is one more seed file, for example `00007_seed_countries_2027.sql`, using
`INSERT ... ON CONFLICT (year, name) DO NOTHING`.

Rules:

- Never edit or rename a migration that has been merged. Add a new one.
- Write migrations that are safe to run on data that already exists (`IF NOT EXISTS`, `ON CONFLICT DO NOTHING`).
- Keep them compatible with the previous version of the service, because old and new instances can run side by side.
- Keep schema changes and seed data in separate files. No secrets or application logic in SQL.

### Run by hand

`make migrate-status`, `make migrate-up` and `make migrate-down` (one step back) use the database from `.env`.
Prefer a new migration over `migrate-down`: rolling back does not restore data that a migration deleted
(for example the de-duplication in `0002`).

### Production

- Take a `pg_dump` before the first start of a version that deletes or rewrites data.
- After a deploy, the startup log lists the applied migrations (`[migrations] applied ...`), or says the database
  is up to date. The applied versions are in the `goose_db_version` table.
- A stuck lock cannot outlive the process: the advisory lock is released when the connection closes.

### Adopting an existing database

A database that was created earlier from the old `initial.sql` needs no manual step. The baseline is a no-op
there, `0002` removes duplicate `score`, `country` and `group` rows (keeping the score row with the most
progress) and adds the keys, and the seeds skip rows that exist. The `group` table is recreated only if
`members` is not already an array, because groups are reset every season.

## Checks

`make check` runs `go vet`, the tests and a scan that no SQL is built with `fmt.Sprintf`.

## Roadmap

- [x] proxy for permission checking
- [x] env var to define admin list
- [x] score calculation

## Docker image

`Dockerfile` is a multi-stage build: dependencies are downloaded in their own cached layer (rebuilds after a
code-only change do not download anything again), the binary is built with `CGO_ENABLED=0 -trimpath -ldflags="-s -w"`,
and the final image is `gcr.io/distroless/static` running as a non-root user with only the binary in it (about 12 MB).
`GO_VERSION` in the Dockerfile must match the `go` line in `go.mod`; CI checks it.

- Build locally: `make docker-build`.
- The image has no shell and no curl. Its `HEALTHCHECK` runs `/app healthcheck`, which requests
  `http://127.0.0.1:$PORT/api/health` and exits 0 on HTTP 200. `/api/health` pings the database, so a database
  outage also makes the container unhealthy.
- `PORT` defaults to `8080` inside the image. The port exposed in Coolify must equal `PORT`.
- Build without make: `go build -trimpath -o bin/app ./cmd/app`; run without building: `go run ./cmd/app`.

## CI and deployment

- **Pull requests** run `.github/workflows/ci.yml`: `gofmt`, `go vet`, `go test ./...` (with a Postgres service, so the
  integration tests run), the Dockerfile/`go.mod` Go version check, and a Docker build. The built image must run as
  non-root, stay under the size limit, migrate an empty database and report healthy. Nothing is pushed.
- **Merging to `main`** runs `.github/workflows/publish.yml`: the same checks, then the image is pushed to
  `ghcr.io/vilmis04/eurovision-game-service` tagged `sha-<commit>` and `latest` (version tags `v*` are also tagged),
  then the Coolify deploy webhook is called.
- **One-time setup**
  - Coolify: create the application from the Docker image `ghcr.io/vilmis04/eurovision-game-service:latest`
    (add registry credentials with `read:packages` if the package is private), set its exposed port to `PORT`, set the
    environment variables from the table above, and leave the UI health check off (the image has no curl; the
    Dockerfile `HEALTHCHECK` is used). Run it next to the old deployment, verify, then move the domain over.
    Do not assign a public domain or published port to the service itself; it is reached through the auth proxy.
  - GitHub: add the repository secrets `COOLIFY_DEPLOY_URL` (the application's deploy webhook URL) and `COOLIFY_TOKEN`
    (a Coolify API token that may deploy it). Without them the deploy step is skipped. Require the `CI` checks
    before merging in the branch protection settings.
  - After the first publish, link the package to this repository in its GitHub package settings.
- **Rollback**: redeploy an older `sha-<commit>` tag.

### Auth proxy

[`vilmis04/auth-proxy`](https://github.com/vilmis04/auth-proxy) must send `X-Internal-Token` with the value of this
service's `INTERNAL_TOKEN` on every request it forwards (for example by setting it on the request in
`ProxyMiddleware` before `proxy.ServeHTTP`). Until it does, every request gets 401.
For local development `docker compose --profile proxy up -d` builds the proxy from GitHub (see `compose.yml` and
`.env.auth.example`).
