FROM golang:1.21.5-alpine as builder

WORKDIR /app

COPY . .

RUN go build .

FROM alpine:latest

WORKDIR /app

COPY --from=builder app/backend-service-go .

CMD [ "./backend-service-go" ]


