# Eurovision guessing game BE service built with Go

## Local setup

### Without database

To run locally just the service: `go run .`

To build an executable: `go build .`

### In docker with db and db adminer tool

To start virtual environment: `make start`

To stop virtual environment: `make stop`

To attach to go-service when container is running: `make attach`
