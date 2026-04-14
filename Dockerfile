# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tarmac .

# Run stage
FROM alpine:latest

RUN apk add --no-cache typst

WORKDIR /app

RUN mkdir -p out/pdf backups

# Copy the built binary
COPY --from=builder /app/tarmac .

EXPOSE 8080

CMD ["./tarmac"]
