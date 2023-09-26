BINARY_NAME := img-dl

all: run

build:
	go build -o ./bin/$(BINARY_NAME) ./cmp/cli/main.go

cross-build:
	GOARCH=amd64 GOOS=linux   go build -o ./bin/$(BINARY_NAME) ./cmp/cli/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/$(BINARY_NAME).exe ./cmp/cli/main.go

run: build
	./bin/$(BINARY_NAME)

clean:
	rm ./bin/*

