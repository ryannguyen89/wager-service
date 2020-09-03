gen-server-models:
	rm -rf server
	openapi-generator-cli generate -i swagger.yml -g go-gin-server -o server \
	--additional-properties=prependFormOrBodyParameters=true,packageName=wager