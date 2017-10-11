OS := $(shell uname)

build: clean
	go build -o bin/bit

test:
	go test -v

clean:
	go clean
	rm -f bit
	rm -f bin/*

darwin:
	GOOS=darwin GOARCH=amd64 go build -v -o bin/bit-darwin

freebsd:
	GOOS=freebsd GOARCH=amd64 go build -v -o bin/bit-freebsd

windows:
	GOOS=windows GOARCH=amd64 go build -v -o bin/bit-windows

linux:
	GOOS=linux GOARCH=amd64 go build -v -o bin/bit-linux

release: clean darwin freebsd windows linux

