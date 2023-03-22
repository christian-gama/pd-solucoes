PWD = $(shell pwd)
GENERATED_DIR ?= $(PWD)/.generated
CACHE_DIR ?= $(PWD)/.cache
BUILD_DIR ?= $(GENERATED_DIR)/build
MAKE = make --no-print-directory


# WORKDIR is used to set the working directory for Dockerfile builds.
export WORKDIR=/go/src/github.com/christian-gama/pg-solucoes

-include $(ENV_FILE)

.PHONY: init
init: cmd-exists-git
	@echo "Initializing git hooks"
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@$(MAKE) remake


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


.PHONY: migrate
migrate:
	@$(MAKE) go run ./cmd/migrate/*.go -e $(ENV_FILE) $(MIGRATION) $(VERSION)


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
	@docker run -v "$(PWD)":/src -w /src alpine:3.17 sh -c "rm -rf $(PWD)/testutils/mocks/*"

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


.PHONY: docker
docker: cmd-exists-docker
	@WORKDIR=$(WORKDIR) docker compose \
		--env-file "$(ENV_FILE)" \
		run \
		-p $(APP_PORT):$(APP_PORT) \
		--rm \
		-e ENV_FILE=$(ENV_FILE) \
		api \
		$(ENTRY_POINT) 


.PHONY: cmd-exists
cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)