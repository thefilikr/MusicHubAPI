FROM docker.io/library/golang:alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN go build -o main ./cmd/app/main.go

EXPOSE 5678
CMD ["goose", "-dir", "./migrations", "postgres", "'postgresql://admin:admin@localhost:5432/song?sslmode=disable'", "up"]
CMD ["./main"]
