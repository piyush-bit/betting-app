# Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

# Pre-copy/cache go.mod for pre-downloading dependencies
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /price-service ./cmd/main.go

EXPOSE 8080

CMD ["/price_service"]
