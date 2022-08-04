.PHONY: build clean deploy

build:
	mkdir bin
	mkdir bin/sqs
	mkdir bin/sqs/ses
	mkdir bin/api
	GOOS=linux go build -o bin/sqs/ses/main sqs/ses/main.go
	GOOS=linux go build -o bin/api/main api/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
