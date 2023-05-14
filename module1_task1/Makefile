# Define variables
POST_TITLE ?= Untitled
POST_NAME ?= untitled
HUGO = hugo
DIST_DIR = dist

# Define targets

## Build the website from markdown and configuration files
build:
    @$(HUGO) --source . --destination $(DIST_DIR) --theme ananke

## Cleanup the contents of the dist directory
clean:
    @rm -rf $(DIST_DIR)/*

## Create a new blog post with a given name and title
post:
    @if [ -z "$(POST_TITLE)" ] || [ -z "$(POST_NAME)" ]; then \
        echo "Please set POST_TITLE and POST_NAME environment variables"; \
        exit 1; \
    fi
    @hugo new --kind post $(DIST_DIR)/$(POST_NAME).md
    @sed -i '' 's/title = ""/title = "$(POST_TITLE)"/' $(DIST_DIR)/$(POST_NAME).md

## Show the available targets and their usage
help:
    @echo "Available targets:"
    @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

