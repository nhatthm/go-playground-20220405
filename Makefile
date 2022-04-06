DATABASE_DSN ?= mongodb://localhost:27017/test

GO ?= go

GO_MIGRATE_VERSION := v4.15.1
MONGODB_VERSION := 4.4

UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
	OS := linux
endif
ifeq ($(UNAME_S),Darwin)
	OS := darwin
endif

UNAME_P := $(shell uname -p)
ifeq ($(UNAME_P),i386)
	ARCH += amd64
endif
ifeq ($(UNAME_P),x86_64)
	ARCH += amd64
endif
ifneq ($(filter arm%,$(UNAME_P)),)
	ARCH += arm64
endif

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

bin/migrate:
	@echo "$(OK_COLOR)==> Install bin/migrate $(NO_COLOR)"
	@test -d bin || mkdir bin
	@curl -sL https://github.com/golang-migrate/migrate/releases/download/$(GO_MIGRATE_VERSION)/migrate.$(OS)-$(ARCH).tar.gz | tar xz migrate \
        && mv migrate bin/migrate

.PHONY: migrations
migrations: bin/migrate
	@echo "$(OK_COLOR)==> Run migrations $(NO_COLOR)"
	@bin/migrate -source=file://./resources/migrations/ -database=$(DATABASE_DSN) up

.PHONY: test-with-docker-compose
test-with-docker-compose: migrations
	@echo "$(OK_COLOR)==> Run test-with-docker-compose $(NO_COLOR)"
	@$(GO) test -gcflags=-l -covermode=atomic -tags=test_with_docker_compose -race -v ./...

.PHONY: test-with-testcontainers
test-with-testcontainers:
	@echo "$(OK_COLOR)==> Run test-with-testcontainers $(NO_COLOR)"
	@$(GO) test -gcflags=-l -covermode=atomic -tags=test_with_testcontainers -race -v ./...

.PHONY: clean
clean:
	@rm -rf bin
