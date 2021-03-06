default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

clean-all:	## Go Clean
	go clean
	rm -rf ./bin/*

test:	## Run Unit Tests
	go test -v ./... -short

test-integration: ## Run Integration TEsts
	go test -v ./...

build: clean-all ## Run dep ensure and build linux binary of all individual functions
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/email functions/email/*.go
