.PHONY: all lint test

all: lint test

lint:
	@go get -u golang.org/x/lint/golint
	golint ./...

test:
	go test -v -cover ./...
