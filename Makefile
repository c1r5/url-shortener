.PHONY: run build deps tidy help deploy


build:
	go build -o bin/url-shortener ./src/main.go

dev:
	@docker compose -f compose.dev.yml down app
	@docker compose -f compose.dev.yml up app --watch --build

database:
	@docker compose -f compose.dev.yml down postgres
	@docker compose -f compose.dev.yml up postgres -d