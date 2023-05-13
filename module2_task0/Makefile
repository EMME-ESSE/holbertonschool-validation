BINARY_NAME=awesome-api
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
BUILD_VERSION=dev
COMMIT=$(shell git rev-parse --short HEAD)

.PHONY: build run stop clean test lint help

build:
	@echo "Compiling $(BINARY_NAME)..."
	go build -ldflags "-s -w -X main.buildTime=$(BUILD_TIME) -X main.buildVersion=$(BUILD_VERSION) -X main.gitCommit=$(COMMIT)" -o $(BINARY_NAME)

run:
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME) >./$(BINARY_NAME).log 2>&1 &

stop:
	@echo "Stopping $(BINARY_NAME)..."
	kill $(shell ps aux | grep '[a]wesome-api' | awk '{print $$2}')

clean:
	@echo "Cleaning $(BINARY_NAME)..."
	-kill $(shell ps aux | grep '[a]wesome-api' | awk '{print $$2}')
	-rm $(BINARY_NAME)
	-rm $(BINARY_NAME).log

test:
	@echo "Testing $(BINARY_NAME)..."
	go test -v ./...

lint:
	@echo "Linting $(BINARY_NAME)..."
	golangci-lint run

help:
	@echo "Project Lifecycle:"
	@echo "  build: compile the source code of the application to a binary named $(BINARY_NAME) using the command 'go build'."
	@echo "  run: run the application in the background by executing the binary $(BINARY_NAME) and write logs into a file named $(BINARY_NAME).log using the command './$(BINARY_NAME) >./$(BINARY_NAME).log 2>&1 &'."
	@echo "  stop: stop the application using the command 'kill XXXXX' where XXXXX is the Process ID of the application."
	@echo "  clean: stop the application and delete the binary $(BINARY_NAME) and the log file $(BINARY_NAME).log."
	@echo "  test: test the application to ensure that it behaves as expected."
	@echo "  lint: run the application through lint to check for errors."
