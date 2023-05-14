build:
	@hugo

clean:
	@rm -rf dist

post:
	@hugo new posts/$(POST_NAME).md
	@echo "POST_TITLE: $(POST_TITLE)"
	@echo "POST_NAME: $(POST_NAME)"
	@echo "---\
	title: '$(POST_TITLE)'\
	date: $(date +"%Y-%m-%dT%H:%M:%S%z")\
	---" > content/posts/$(POST_NAME).md