BINARY_NAME := img-dl

all: run

build:
	go build -o ./bin/$(BINARY_NAME) ./cmp/app/main.go

cross-build:
	GOARCH=amd64 GOOS=linux   go build -o ./bin/$(BINARY_NAME) ./cmp/app/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/$(BINARY_NAME).exe ./cmp/app/main.go

run: build
	./bin/$(BINARY_NAME)

clean:
	rm ./bin/*

