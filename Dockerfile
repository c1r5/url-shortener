# build stage
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

# final stage
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/bin/url-shortener /app/url-shortener

EXPOSE 3001
RUN chmod +x /app/url-shortener
CMD ["/app/url-shortener"]