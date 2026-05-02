.PHONY: run build deps tidy help deploy

GOOS ?= linux
GOARCH ?= amd64

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/server ./src/main.go

dev: build
	@docker compose -f compose.dev.yml down app
	@docker compose -f compose.dev.yml up app --watch --build

database:
	@docker compose -f compose.dev.yml down postgres
	@docker compose -f compose.dev.yml up postgres -d

DEPLOY_DIR ?= ~/url-shortener

deploy:
	DEPLOY_DIR=$(DEPLOY_DIR) ./deploy.sh
