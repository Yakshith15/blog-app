FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog-service cmd/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/blog-service /app/blog-service

EXPOSE 8082

ENTRYPOINT ["/app/blog-service"]
