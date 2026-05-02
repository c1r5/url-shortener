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
COPY --from=builder /app/bin/server /app/server

EXPOSE 3001
RUN chmod +x /app/server
CMD ["/app/server"]
