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
	tar -czvf bin/bit-darwin.tar.gz bin/bit-darwin
freebsd:
	GOOS=freebsd GOARCH=amd64 go build -v -o bin/bit-freebsd
	tar -czvf bin/bit-freebsd.tar.gz bin/bit-freebsd

windows:
	GOOS=windows GOARCH=amd64 go build -v -o bin/bit-windows
	tar -czvf bin/bit-windows.tar.gz bin/bit-windows

linux:
	GOOS=linux GOARCH=amd64 go build -v -o bin/bit-linux
	tar -czvf bin/bit-linux.tar.gz bin/bit-linux

release: clean darwin freebsd windows linux

