start:
		docker compose up -d --build && docker exec -it go-service sh

stop:
		docker compose stop

restart:
		docker compose stop && docker compose up -d --build && docker exec -it go-service sh

attach:
		docker exec -it go-service sh

remove:
		docker compose down

build:
		go build ./cmd/app -o /bin

# Loads .env into the environment (the service itself no longer reads .env files).
run:
		set -a; . ./.env; set +a; go run ./cmd/app

# Migrations are applied automatically when the service starts. These targets are
# for inspecting or stepping them by hand against the database in .env.
GOOSE = go run github.com/pressly/goose/v3/cmd/goose@v3.26.0 -dir internal/migrations postgres "$$DATABASE_URL?sslmode=$${DB_SSLMODE:-disable}"

migrate-status:
		set -a; . ./.env; set +a; $(GOOSE) status

migrate-up:
		set -a; . ./.env; set +a; $(GOOSE) up

# Rolls back one migration. Data deleted by a migration is not restored.
migrate-down:
		set -a; . ./.env; set +a; $(GOOSE) down

# make migrate-new name=add_something
migrate-new:
		@test -n "$(name)" || (echo "usage: make migrate-new name=add_something" && exit 1)
		go run github.com/pressly/goose/v3/cmd/goose@v3.26.0 -dir internal/migrations -s create $(name) sql

# Needs a Postgres superuser URL, e.g. postgres://user:pass@localhost:5432/postgres?sslmode=disable
test-integration:
		@test -n "$$TEST_DATABASE_URL" || (echo "set TEST_DATABASE_URL" && exit 1)
		go test ./...

check:
		go vet ./...
		go test ./...
		@! grep -rn --include=*.go --exclude=*_test.go -E 'Sprintf\(.*(SELECT|INSERT|UPDATE|DELETE)' internal cmd || (echo "SQL built with Sprintf" && exit 1)
