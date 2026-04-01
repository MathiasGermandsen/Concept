# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/server .

# Runtime stage
FROM alpine:3.21

RUN apk --no-cache add ca-certificates \
    && addgroup -S appgroup \
    && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /build/server .

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

ENTRYPOINT ["./server"]
