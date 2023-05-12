Makefile for Awesome API
This Makefile is designed to automate the life-cycle of Awesome API. It provides several targets to build, run, stop, clean and test the application.

Targets
build: compiles the application.
run: starts the application and writes output to awesome-api.log.
stop: stops the application.
clean: stops the application and removes the binary and log files.
test: sends HTTP requests to the running application for testing purposes.
help: prints a list of available targets with descriptions.
Usage
To build the application, run the following command:


make build
To run the application, execute:


make run
To stop the application, execute:


make stop
To clean the binary and log files, run:


make clean
To test the application, run:


make test
To view a list of available targets with descriptions, run:


make help
