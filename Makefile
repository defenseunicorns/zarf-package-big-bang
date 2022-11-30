.DEFAULT_GOAL := help

# Idiomatic way to force a target to always run, by having it depend on this dummy target
FORCE:

.PHONY: help
help: ## Show a list of all targets
	@grep -E '^\S*:.*##.*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1:\3/p' \
	| column -t -s ":"

.PHONY: clean
clean: ## Clean up build files
	@rm -rf ./build zarf-sbom

.PHONY: build
build: ## Build the Big Bang Zarf package
	./scripts/build.sh

.PHONY: test
test: ## Run e2e test of the package
	./scripts/tests.sh
