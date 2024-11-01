# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

## Install dependencies for Alpine Linux
RUN apk add --no-cache git

## Set the working directory inside the builder stage
WORKDIR /app

## Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

## Copy the source code
COPY . .

## Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o golang-api .

# Stage 2: Create a minimal runtime environment
FROM alpine:latest

## Set the working directory inside the runtime container
WORKDIR /app

## Copy the built binary from the builder stage
COPY --from=builder /app/golang-api .

## Expose port for Gin server
EXPOSE 5000

## Run the application
ENTRYPOINT ["/app/golang-api"]
