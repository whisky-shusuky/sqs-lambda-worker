.PHONY: build clean deploy

build:
	mkdir bin
	mkdir bin/sqs
	mkdir bin/sqs/ses
	GOOS=linux go build -o bin/sqs/ses/main sqs/ses/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
