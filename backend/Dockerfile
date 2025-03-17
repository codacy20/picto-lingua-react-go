FROM golang:1.24-alpine AS builder

WORKDIR /app

# Add git and certificates for private repositories
RUN apk add --no-cache git ca-certificates

# Set environment variables to skip GOPROXY verification
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=off

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /picto-lingua-backend

# Create a minimal image
FROM alpine:latest

WORKDIR /app

# Add ca-certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Copy binary from builder stage
COPY --from=builder /picto-lingua-backend .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./picto-lingua-backend"] 