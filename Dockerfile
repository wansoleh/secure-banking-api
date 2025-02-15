# Menggunakan base image Go
FROM golang:1.19

# Set working directory dalam container
WORKDIR /app

# Copy go.mod dan go.sum untuk dependency management
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

# Copy semua source code ke dalam container
COPY . .

# Build aplikasi
RUN go build -o banking-api

# Expose port yang digunakan
EXPOSE 8080

# Jalankan aplikasi
CMD ["/app/banking-api"]
