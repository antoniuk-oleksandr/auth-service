# Build Stage
FROM golang:1.25.1-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/app/main.go

# Final Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=0 /app/main .

EXPOSE 8080
CMD ["./main"]