FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o banking-api main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/banking-api .
EXPOSE 8080
CMD ["./banking-api"]
