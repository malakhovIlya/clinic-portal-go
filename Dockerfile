# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/clinic/main.go

# Runtime stage
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/server ./server
COPY --from=builder /app/web/dist ./web/dist
COPY --from=builder /app/web/public ./web/public

RUN adduser -D appuser
USER appuser

EXPOSE 8080
ENTRYPOINT ["./server"]