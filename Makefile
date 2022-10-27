# The version of Big Bang to use.
BIGBANG_VERSION := 1.43.0

# The version of Zarf to use. To keep this repo as portable as possible the Zarf binary will be downloaded and added to
# the build folder.
ZARF_VERSION := v0.21.3

# Figure out which Zarf binary we should use based on the operating system we are on
ZARF_BIN := zarf

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
	@rm -rf ./build

.PHONY: all
all: | build

mkdir:
	@mkdir -p build

build: mkdir ## Build the Big Bang Zarf Package 
	@echo "Creating the deploy package"
	@$(ZARF_BIN) package create --skip-sbom --confirm
	@mv zarf-package-big-bang-amd64.tar.zst build/zarf-package-big-bang-amd64.tar.zst
