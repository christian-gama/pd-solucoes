init:
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@echo "Initialized git hooks."

tidy:
	@go mod tidy
	@go mod vendor