.DEFAULT_GOAL := start
.PHONY: fetch help start

help: ## Print this help message
	@echo "List of available make commands";
	@echo "";
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}';
	@echo "";

fetch: ## ?
	go run . --depth 2 --fetch

install: ## Build and install the application
	go build
	go install

start: ## ?
	go run . --depth 2 --no-root
