FROM golang:latest AS builder

WORKDIR /app

# Download and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build the app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Create runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
