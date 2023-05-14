build:
	@hugo

clean:
	@rm -rf dist

post:
	@POST_NAME=$$(openssl rand -hex 1); \
	POST_NAME+= $$(openssl rand -hex 1); \
	POST_TITLE=$$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1); \
	hugo new posts/$$POST_NAME.md; \
	echo "---\n\
	title: '$$POST_TITLE'\n\
	date: $$(date +"%Y-%m-%dT%H:%M:%S%z")\n\
	---" > content/posts/$$POST_NAME.md
