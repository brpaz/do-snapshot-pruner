PROJECT_NAME := "do-snapshot-pruner"

.SHELLFLAGS = -c # Run commands in a -c flag
.ONESHELL: ; # recipes execute in same shell
.NOTPARALLEL: ; # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.DEFAULT_GOAL:=help
.PHONY: fmt lint vet test test-cover build-docker release help

fmt: ## Runs code formatter
	@gofmt -l -w .

lint: ## Lint code
	@docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.25.0 golangci-lint run -v

test: ## Runs tests
	@go test -v ./...

test-cover: ## Runs tests and generate coverage report
	@mkdir -p coverage
	@go test -v -coverprofile coverage/cover.out ./...
	@go tool cover -html=coverage/cover.out -o coverage/cover.html

build-docker: ## Builds a Docker image for the project
	@docker build \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	--build-arg VCS_REF="" \
	--build-arg VERSION=0.0.0-dev . -t $(PROJECT_NAME)

release-check: ## Go release validation
	@goreleaser --snapshot --skip-publish --rm-dist

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
