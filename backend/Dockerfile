# Build Stage
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get github.com/antoniuk-oleksandr/auth-service/ctp
RUN go get github.com/antoniuk-oleksandr/auth-service/common

RUN go build -o main ./cmd/auth-backend/main.go

# Final Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
