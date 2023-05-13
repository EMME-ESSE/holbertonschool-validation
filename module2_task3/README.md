# Golang Project

## Prerequisites

goland
make

## Lifecycle

lint: Static linting on go files using golangci-lint

build:  Compile the source code of the application to a binary named awesome-api

run:  Run the application in background by executing the binary awesome-api, and write logs into a file named awesome-api.log

stop:  Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application. 

clean:  Delete the binary awesome-api and the log file awesome-api.log

test:  Check expected behaviour of the app

help:  Show this help message

unit-tests: Execute (successfully) the Golang unit tests


integration-tests: Execute (successfully) the Golang integration tests