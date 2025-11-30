FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

COPY . .

RUN go build -o scraper-service .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/scraper-service .

CMD ["./scraper-service"]
