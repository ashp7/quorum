.PHONY: all build test coverage lint complexity clean

BINARY_NAME := contractclient

all: deps build test coverage complexity

deps:
	go get github.com/fzipp/gocyclo

build:
	go fmt ./
	go vet ./
	go build -o $(BINARY_NAME) *.go

test:
	go test -v ./

coverage:
	go test -coverprofile=coverage.out ./
	go tool cover -html=coverage.out -o coverage.html

complexity:
	gocyclo *.go 

clean:
	go clean
	rm -f $(BINARY_NAME) coverage.out coverage.html

