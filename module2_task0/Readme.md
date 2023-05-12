Awesome API
This is a simple API application that can be built, run, tested, and stopped with a Makefile.

Prerequisites
To run this application, you must have the following software installed on your system:

go (v1.16 or higher)
curl
make
Usage
To build the application, run:


make build
To run the application, run:

make run
This will start the application and redirect its output to ./awesome-api.log. To stop the application, run:


make stop
To clean up the application and its log file, run:

make clean
To test the application, run:

make test
This will send requests to http://localhost:9999 and http://localhost:9999/health using curl.

Makefile Commands
Here is a list of available Makefile commands:

build: Build the awesome-api binary.
run: Run the awesome-api application.
stop: Stop the awesome-api application.
clean: Stop and clean the awesome-api application and log.
test: Test the awesome-api application by sending requests to http://localhost:9999 and http://localhost:9999/health.
help: Display what each command does.
