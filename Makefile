all: clean lint test

test:
	go test -v ./...

lint: testdeps
	go vet ./...
	golint -set_exit_status ./...

testdeps:
	go get -d -v -t ./...
	go get golang.org/x/lint/golint

clean:
	go clean

.PHONY: test lint testdeps clean
