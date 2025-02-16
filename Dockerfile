# Use Go 1.23.4
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Install git (needed for some Go dependencies)
RUN apk add --no-cache git

# Copy Go modules first for caching
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Copy the rest of the app
COPY . .

# Build the Go application
RUN go build -o banking-api .

# Use a smaller final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/banking-api .

# Expose port 8080
EXPOSE 8080

# Start the app
CMD ["./banking-api"]
