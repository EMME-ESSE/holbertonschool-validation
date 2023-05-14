PUBLIC = dist
RM = rm -r

.PHONY: clean
clean: ; $(RM) ./${PUBLIC}/*

.PHONY: build
build:
	@hugo -d $(PUBLIC)

.PHONY: post
post:; 
	@hugo new posts/$(POST_NAME).md
	@sed -i '/title:/c\title: "$(POST_TITLE)"' content/posts/$(POST_NAME).md
