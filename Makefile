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

run:
		go run ./cmd/app

check:
		go vet ./...
		go test ./...
		@! grep -rn --include=*.go --exclude=*_test.go -E 'Sprintf\(.*(SELECT|INSERT|UPDATE|DELETE)' internal cmd || (echo "SQL built with Sprintf" && exit 1)
