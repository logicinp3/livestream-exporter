# todo: build Dockerfile and push
#

.PHONY: test
test: ## Test step
	@echo test

linux:
	GOOS=linux GOARCH=amd64 go build
mac:
	GOOS=darwin GOARCH=amd64 go build

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
