PWD = $(shell pwd)
GENERATED_DIR ?= $(PWD)/.generated
CACHE_DIR ?= $(PWD)/.cache
BUILD_DIR ?= $(GENERATED_DIR)/build
MAKE = make --no-print-directory
APP_NAME = pd-solucoes

# WORKDIR is used to set the working directory for Dockerfile builds.
export WORKDIR=/go/src/github.com/christian-gama/pd-solucoes

-include $(ENV_FILE)

.PHONY: init
init: cmd-exists-git
	@echo "Initializing git hooks"
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@chmod +x ./scripts/*.sh
	@./scripts/create-env.sh
	@$(MAKE) remake


.PHONY: run
run: cmd-exists-go clear-screen cmd-exists-gin
ifeq ($(ENV_FILE), .env.prod)
	@$(MAKE) build
	@$(BUILD_DIR)/$(APP_NAME) -e $(ENV_FILE)
else
	@gin --port 3000 --appPort $(APP_PORT) --build ./cmd/api --path . --bin $(APP_NAME) -i run -e $(ENV_FILE)
endif


.PHONY: build
build: cmd-exists-go
	@echo "Generating build for $(APP_NAME) using $(ARCH) architecture..."
	@CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/*.go
	@echo "Build was generated at $(BUILD_DIR)/$(APP_NAME)"


.PHONE: tidy
tidy:
	@go mod tidy
	@go mod vendor


.PHONY: clear-screen
clear-screen:
	@printf "\033c"


.PHONY: migrate-
migrate-%: cmd-exists-docker
	@if [ -z "$(*)" ]; then \
		echo "Error: expected [up|down|force|drop|steps|version]"; \
		exit 1; \
	fi;
	@MIGRATION=$(*) $(MAKE) migrate


.PHONY: migrate-
migrate-%: cmd-exists-docker
	@if [ -z "$(*)" ]; then \
		echo "Error: expected [up|down|force|drop|steps|version]"; \
		exit 1; \
	fi;
	@ENV_FILE=.env.dev MIGRATION=$(*) $(MAKE) migrate


.PHONY: migrate
migrate: cmd-exists-docker
	@$(MAKE) docker ENV_FILE=$(ENV_FILE) ENTRY_POINT="go run ./cmd/migrate/*.go -e $(ENV_FILE) $(MIGRATION) $(VERSION)"


.PHONY: psql-open
psql-open: cmd-exists-docker
	@if [ -z "$(ENV_FILE)" ]; then \
		echo "Error: expected ENV_FILE"; \
		exit 1; \
	fi;

	@if [ "$(ENV_FILE)" != ".env.dev" ] && [ "$(ENV_FILE)" != ".env.prod" ]; then \
		echo "Error: expected .env.dev or .env.prod"; \
		exit 1; \
	fi;

	@docker compose --env-file $(ENV_FILE) up -d psql


.PHONY: create-migration
create-migration: cmd-exists-docker
	@if [ -z "$(NAME)" ]; then \
		echo "Error: expected NAME"; \
		exit 1; \
	fi;
	@docker run -it \
		-v $(PWD)/migrations:/migrations \
		--rm migrate/migrate \
		create -ext sql -dir ./migrations $(NAME)


.PHONY: test-unit
test-unit: cmd-exists-go clear-screen
	@TEST_MODE=unit go test ./...


.PHONY: test-integration
test-integration: cmd-exists-go clear-screen
	@docker compose --env-file .env.test up -d psql_test
	@go run ./cmd/migrate/*.go -e .env.test drop ""
	@go run ./cmd/migrate/*.go -e .env.test up ""
	@./scripts/wait-for-db.sh pd-solucoes-psql-test TEST_MODE=integration go test ./...


.PHONY: test
test: cmd-exists-go clear-screen
	@docker compose --env-file .env.test up -d psql_test
	@go run ./cmd/migrate/*.go -e .env.test drop ""
	@go run ./cmd/migrate/*.go -e .env.test up ""
	@./scripts/wait-for-db.sh pd-solucoes-psql-test TEST_MODE=integration go test ./...


.PHONY: lint
lint: cmd-exists-docker
	@sh -c "./scripts/linter.sh"

	
.PHONY: mock
mock: cmd-exists-docker
	@echo "Generating mocks..."
	@docker run -v "$(PWD)":/src -w /src alpine:3.17 sh -c "rm -rf testutils/mocks/*"

	@docker run \
		-v "$(PWD)":/src \
		-w /src vektra/mockery --all --keeptree --case underscore --exported --dir ./internal --quiet --output ./testutils/mocks

	@docker run \
		-v "$(PWD)":/src \
		-w /src vektra/mockery --all --keeptree --case underscore --exported --dir ./pkg --quiet --output ./testutils/mocks
	@echo "Mocks generated successfully"


.PHONY: clear
clear:
	@echo "Clearing..."
	@docker run -v "$(PWD)":/src -w /src alpine:3.17 sh -c "rm -rf \
		testutils/mocks \
		.generated/ \
		.cache/"
	@echo "Clear done"


.PHONY: remake
remake:
	@$(MAKE) clear
	@$(MAKE) mock
	@go mod tidy
	@go mod vendor
	@$(MAKE) migrate-drop ENV_FILE=.env.dev
	@$(MAKE) migrate-up ENV_FILE=.env.dev


.PHONY: docker
docker: cmd-exists-docker
	@WORKDIR=$(WORKDIR) APP_NAME=$(APP_NAME) docker compose \
		--env-file "$(ENV_FILE)" \
		run \
		--name $(APP_NAME) \
		-p $(APP_PORT):$(APP_PORT) \
		--rm \
		-e ENV_FILE=$(ENV_FILE) \
		--build \
		api \
		$(ENTRY_POINT) 


.PHONY: docker-down
docker-down: cmd-exists-docker
	@docker compose --env-file ".env.dev" down 


.PHONY: docker-rebuild
docker-rebuild: cmd-exists-docker
	@docker compose --env-file ".env.dev" build


.PHONY: docker-dev
docker-dev: cmd-exists-docker
	@$(MAKE) docker ENV_FILE=.env.dev ENTRY_POINT="make run"


.PHONY: docker-prod
docker-prod: cmd-exists-docker
	@$(MAKE) docker ENV_FILE=.env.prod ENTRY_POINT="make run" 


.PHONY: cmd-exists
cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)