# Stage 1: Build the application
FROM golang:1.27-alpine AS builder

# Add the author
LABEL authors="franklinesable"

WORKDIR /app

# Copy module file
COPY go.mod ./

# Copy application source
COPY . .

# Compile Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o shortener .

# Stage 2: Run the application
FROM alpine:3.22

WORKDIR /app

# Copy executable
COPY --from=builder /app/shortener .

# Copy frontend files
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/styles ./styles

EXPOSE 8080

CMD ["./shortener"]