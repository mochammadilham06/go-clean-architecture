# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod dan sum files
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary (sesuaikan path main.go anda)
# CGO_ENABLED=0 agar binary static (bisa jalan di alpine polos)
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Stage 2: Run (Image kecil)
FROM alpine:latest

WORKDIR /root/

# Copy binary dari stage builder
COPY --from=builder /app/myapp .
# Copy .env jika perlu (tapi nanti di prod pakai environment variable cloud)
# COPY .env . 

# Expose port (sesuai config)
EXPOSE 9091

# Jalankan app
CMD ["./myapp"]