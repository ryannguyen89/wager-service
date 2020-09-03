gen-server-models:
	rm -rf server
	openapi-generator-cli generate -i swagger.yml -g go-gin-server -o server \
	--additional-properties=prependFormOrBodyParameters=true,packageName=wager

build:
	GOOS=linux CGO_ENABLED=0 go build -o wager-service main.go

run: build
	./wager-service
