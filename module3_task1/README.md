make check
	hugo

clean:
	-rm -rf ./dist
	-make stop
	-rm ./awesome-api ./awesome-api.log ./coverage-units.out

post:
	hugo new posts/${POST_NAME}.md
	sed -i 's/^title:.*/title: ${POST_TITLE}/g' ./content/posts/${POST_NAME}.md

run:
	./awesome-api >./awesome-api.log 2>&1 &

stop:
	kill "$(shell pgrep awesome-api)"

test:
	make unit-tests
	make integration-tests

lint:
	@golangci-lint run

unit-tests:
	@go test -v -short -coverprofile=coverage-units.out

integration-tests:
	@go test -v -coverprofile=coverage-integrations.out

check:
	markdown-link-check ./content/posts/*
	markdownlint ./content/posts/*

validate:
	- ./w3c_validator.py ./dist/index.html

build: ## Generate the website from the markdown and configuration files in the directory dist/ and compile the source code of the application to a binary named awesome-api
	make check
	hugo
	go build -o awesome-api .

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
