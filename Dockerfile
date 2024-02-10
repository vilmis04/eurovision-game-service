FROM golang:1.21.5-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o /bin/eurovision-game-service ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder app/bin/eurovision-game-service .

CMD [ "./eurovision-game-service" ]


