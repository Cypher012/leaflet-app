include .env
export


# Tools
GO_BIN := $(shell which go)
TEMPL_BIN := $(shell which templ)
AIR_BIN := $(shell which air)

# Project info
PROJECT_NAME := $(shell basename $(CURDIR))

# DB info
DB_DRIVER := postgres
DB_DIR := internal/platform/postgres/migrations
DB_URL := $(DATABASE_URL)
# Build configuration
BUILD_DIR := build
MAIN_PKG := ./cmd/$(PROJECT_NAME)/...

# Simple color scheme
CYAN := \033[36m
DIM := \033[2m
RESET := \033[0m

# Symbols
CHECK := ✓
ARROW := →

.PHONY: check-deps init create-dirs setup-go build serve watch clean help templ migrate-up migrate-down swag-init feat

check-deps:
	@printf "\n$(CYAN)Checking dependencies$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@test -n "$(GO_BIN)" || (printf "✗ Go not installed\n" && exit 1)
	@test -n "$(TEMPL_BIN)" || (printf "✗ templ not installed\n" && exit 1)
	@test -n "$(AIR_BIN)" || (printf "✗ air not installed\n" && exit 1)
	@printf "$(CHECK) All dependencies ready\n"

create-dirs:
	@printf "\n$(CYAN)Creating project structure$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"

	# cmd
	@mkdir -p cmd/$(PROJECT_NAME)

	# views
	@mkdir -p views/{components,layouts,pages}

	# internal (reusable core only)
	@mkdir -p internal/module
	@mkdir -p internal/shared/{middleware,utils,response,validator}
	@mkdir -p internal/platform/{db,postgres}

	# build dir
	@mkdir -p $(BUILD_DIR)

	# main entry
	@[ -f cmd/$(PROJECT_NAME)/main.go ] || printf 'package main\n\nfunc main() {\n\t// TODO: Initialize your server here\n}\n' > cmd/$(PROJECT_NAME)/main.go


	# shared layer
	@[ -f internal/shared/middleware/middleware.go ] || echo "package middleware" > internal/shared/middleware/middleware.go
	@[ -f internal/shared/utils/utils.go ] || echo "package utils" > internal/shared/utils/utils.go
	@[ -f internal/shared/response/response.go ] || echo "package response" > internal/shared/response/response.go
	@[ -f internal/shared/validator/validator.go ] || echo "package validator" > internal/shared/validator/validator.go

	# platform layer
	@[ -f internal/platform/db/db.go ] || echo "package db" > internal/platform/db/db.go
	@[ -f internal/platform/postgres/conn.go ] || echo "package postgres" > internal/platform/postgres/conn.go

	# views
	@[ -f views/layouts/BaseLayout.templ ] || echo "// Base layout component" > views/layouts/BaseLayout.templ
	@[ -f views/components/Navbar.templ ] || echo "// Navigation component" > views/components/Navbar.templ
	@[ -f views/pages/HomePage.templ ] || echo "// Home page component" > views/pages/HomePage.templ

	@printf "$(CHECK) Project structure created\n"

feat:
	@name=$(filter-out $@,$(MAKECMDGOALS)); \
	if [ -z "$$name" ]; then \
		name="$(name)"; \
	fi; \
	if [ -z "$$name" ]; then \
		printf "Usage: make feat <name> OR make feat name=<name>\n"; \
		exit 1; \
	fi; \
	printf "\n$(CYAN)Creating feature: $$name$(RESET)\n"; \
	printf "$(DIM)────────────────────────────────────$(RESET)\n"; \
	mkdir -p internal/module/$$name; \
	[ -f internal/module/$$name/route.go ] || echo "package $$name" > internal/module/$$name/route.go; \
	[ -f internal/module/$$name/handler.go ] || echo "package $$name" > internal/module/$$name/handler.go; \
	[ -f internal/module/$$name/service.go ] || echo "package $$name" > internal/module/$$name/service.go; \
	[ -f internal/module/$$name/repository.go ] || echo "package $$name" > internal/module/$$name/repository.go; \
	[ -f internal/module/$$name/model.go ] || echo "package $$name" > internal/module/$$name/model.go; \
	printf "$(CHECK) Feature '$$name' created\n"

%:
	@:

setup-go:
	@printf "\n$(CYAN)Initializing Go module$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(GO_BIN) mod init $(PROJECT_NAME) 2>/dev/null || true
	@$(GO_BIN) mod tidy >/dev/null 2>&1
	@printf "$(CHECK) Go modules ready\n"

air-init:
	@printf "\n$(CYAN)Initializing Air$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@printf '[build]\n  bin = "./tmp/main"\n  cmd = "go build -o ./tmp/main ./cmd/$(PROJECT_NAME)/."\n  include_ext = ["go", "templ"]\n  exclude_regex = ["_test.go", "_templ.go"]\n  exclude_dir = ["tmp", "vendor"]\n  pre_cmd = ["templ generate"]\n  delay = 1000\n' > .air.toml
	@printf "$(CHECK) Air initialized\n"

templ:
	@printf "\n$(CYAN)Generating templ files$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(TEMPL_BIN) generate
	@printf "$(CHECK) Templ files generated\n"

build: check-deps templ
	@printf "\n$(CYAN)Building project$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(GO_BIN) build -o $(BUILD_DIR)/$(PROJECT_NAME) $(MAIN_PKG)
	@printf "$(CHECK) Build complete\n"

serve: build
	@printf "\n$(CHECK) Starting server...\n"
	@./$(BUILD_DIR)/$(PROJECT_NAME)

watch:
	@printf "\n$(CYAN)Watching for changes$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(AIR_BIN)

clean:
	@printf "\n$(CYAN)Cleaning build files$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@rm -rf $(BUILD_DIR)
	@find . -type f -name "*_templ.go" -delete
	@printf "$(CHECK) Clean complete\n"

init: check-deps create-dirs setup-go air-init
	@$(MAKE) feat name=auth
	@printf "\n$(CHECK) Project setup complete!\n"

test:
	@printf "\n$(CYAN)Running tests$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(GO_BIN) test -v -count=1 -cover ./...

# --- DB migrations ---

migrate-up:
	goose -dir $(DB_DIR) $(DB_DRIVER) $(DB_URL) up

migrate-down:
	goose -dir $(DB_DIR) $(DB_DRIVER) $(DB_URL) down

migrate-reset:
	goose -dir $(DB_DIR) $(DB_DRIVER) $(DB_URL) reset

migrate-status:
	goose -dir $(DB_DIR) $(DB_DRIVER) $(DB_URL) status

migrate-create:
	goose -dir $(DB_DIR) create $(name) sql

swag-init:
	swag init -g main.go -d cmd/server,internal/module/auth,internal/module/feed,internal/module/comment,internal/module/like,internal/module/profile,internal/module/upload,internal/shared/types,internal/shared/response \
	--parseInternal --parseDependency

help:
	@printf "\n$(CYAN)Available commands$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@printf "  make init$(RESET)      $(ARROW) Initialize project\n"
	@printf "  make build$(RESET)     $(ARROW) Build project\n"
	@printf "  make serve$(RESET)     $(ARROW) Run server\n"
	@printf "  make watch$(RESET)     $(ARROW) Hot reload (Air)\n"
	@printf "  make clean$(RESET)     $(ARROW) Clean build\n"
	@printf "  make test$(RESET)      $(ARROW) Run tests\n\n"

	@printf "  make swag-init$(RESET)      $(ARROW) Generate Swagger docs\n\n"

	@printf "  make migrate-up$(RESET)    $(ARROW) Apply migrations\n"
	@printf "  make migrate-down$(RESET)  $(ARROW) Rollback migration\n"
	@printf "  make migrate-reset$(RESET) $(ARROW) Reset DB\n"
	@printf "  make migrate-status$(RESET)$(ARROW) Migration status\n"
	@printf "  make migrate-create$(RESET)$(ARROW) Create migration\n\n"