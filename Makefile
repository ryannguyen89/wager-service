gen-server-models:
	rm -rf server
	openapi-generator-cli generate -i swagger/swagger.yml -g go-gin-server -o server \
	--additional-properties=prependFormOrBodyParameters=true,packageName=wager

gen-client:
	rm -rf clients/wager/client
	openapi-generator-cli generate -i swagger/swagger.yml -g go -o clients/wager/client \
	--additional-properties=prependFormOrBodyParameters=true,packageName=wager

build:
	GOOS=linux CGO_ENABLED=0 go build -o wager-service main.go

run: build
	./wager-service

test:
	go test

docker:
	chmod +x entry-point.sh
	docker build . -t wager
