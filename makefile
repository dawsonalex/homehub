#BINARY ?= $(shell basename "$(PWD)")# binary name
SERVICE_NAME := $(subst cmd/,,$(wildcard cmd/*))
CMD := $(wildcard cmd/*/*.go)
PROTOC_FILE := $(wildcard cmd/*/schema/*.proto)
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

# Clean the build directory (before committing code, for example)
.PHONY: clean
clean: 
	rm -rv bin

PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 darwin/arm64

release: $(PLATFORMS)

build-protoc: $(PROTOC_FILE)

$(PROTOC_FILE):
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
        $@

.PHONY: build-protoc $(PROTOC_FILE)

$(PLATFORMS): build-protoc
	GOOS=$(os) GOARCH=$(arch) go build -o 'bin/$(os)-$(arch)/$(SERVICE_NAME)' $(CMD)

.PHONY: release $(PLATFORMS)

include infrastructure/local/.env

GOOSE_RUN := go run github.com/pressly/goose/cmd/goose@latest
JET_RUN:= go run github.com/go-jet/jet/v2/cmd/jet@latest
POSTGRES_URI = "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)"

.PHONY:up
up:
	docker compose -f ./infrastructure/local/docker-compose.yaml up

.PHONY: add-migration
add-migration:
	$(GOOSE_RUN) -dir db/migrations postgres $(POSTGRES_URI) create new-migration sql

.PHONY: add-fixture
add-fixture:
	$(GOOSE_RUN) -dir db/fixtures postgres $(POSTGRES_URI) create new-fixture sql

.PHONY:migrate-up
migrate-up:
	$(GOOSE_RUN) -dir db/migrations postgres $(POSTGRES_URI) up

.PHONY:migrate-down
migrate-down:
	$(GOOSE_RUN) -dir db/migrations postgres $(POSTGRES_URI) down

.PHONY:fixtures-up
fixtures-up:
	$(GOOSE_RUN) -dir db/fixtures postgres $(POSTGRES_URI) up

.PHONY:gen-db
gen-db:
	$(JET_RUN) -dsn=$(POSTGRES_URI) -schema=public -ignore-tables goose_db_version -path=pkg/db/postgres/gen