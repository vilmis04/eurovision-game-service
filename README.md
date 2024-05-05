# Eurovision guessing game BE service built with Go

## Local setup

### Without database

To run locally just the service: `go run .`

To build an executable: `go build .`

### In docker with db and db adminer tool

To start virtual environment: `make start`

To stop virtual environment: `make stop`

To attach to go-service when container is running: `make attach`

## Roadmap

- [ ] proxy for permission checking
- [ ] env var to define admin list
- [x] score calculation

## Deployment

- run command to deploy to image to dockerhub (replace x.x.x with version): `docker build -t vsud/ev-game:service-x.x.x . && docker push vsud/ev-game:service-x.x.x`
- ssh into the server, update docker compose with the new image version tag: `vi ~/PROJECTS/docker-compose.yaml`
- run `docker compose up -d` to start the service with the changes
