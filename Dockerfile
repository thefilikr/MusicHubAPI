FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main /app/cmd/app/app ./cmd/app/main.go

EXPOSE 8080
CMD ["/app/cmd/app/app"]
