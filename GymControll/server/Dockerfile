FROM golang:alpine AS builder
WORKDIR /app
COPY . .
WORKDIR /app/cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server .

CMD ["/app/server"]