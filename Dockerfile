FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download || true

COPY cmd/ ./cmd/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-world ./cmd/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/hello-world .

EXPOSE 8080

CMD ["./hello-world"]
