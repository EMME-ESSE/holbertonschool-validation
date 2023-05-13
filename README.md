# Project Name

This project provides a set of commands to build and manage a website with Hugo. The website can be created, posts can be added and the site can be cleaned up.

## Prerequisites

Before running this project, make sure the following prerequisites are met:
- Hugo should be installed in the environment where the website will be built and run.
- Basic understanding of command-line interface.

## Project Lifecycle

The following commands are available:

### Build
Command to build the website using Hugo.

```
make build
```

### Clean
Command to clean the files in the `dist` folder.

```
make clean
```

### Post
Command to create a new post with the given `POST_NAME` and `POST_TITLE`. It uses Hugo's `new` command to create a new post and then modifies the `title` and `draft` parameters.

```
make post POST_NAME=<post-name> POST_TITLE="<post-title>"
```

## Requirements

- A Makefile should be present and valid.
- The website should be created with Hugo.
- The `build`, `clean`, and `post` goals should be implemented and mapped to the lifecycle stages of the same name.
- The `README.md` file should exist.

## Usage
To use this project, simply clone the repository and run the desired command using `make` in the project directory.