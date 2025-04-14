FROM golang:1.22.2-alpine AS builder

RUN apk add --no-cache git

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.24.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /www/

ENV GIN_MODE=release

WORKDIR /app

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/.env .

COPY --from=builder /app/migrations ./migrations

COPY --from=builder /app/main .

COPY entrypoint.sh .

RUN chmod +x ./entrypoint.sh

CMD ["./entrypoint.sh"]
