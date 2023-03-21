init:
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@echo "Initialized git hooks."
