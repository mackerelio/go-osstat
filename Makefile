.PHONY: all
all: clean lint test

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint: testdeps
	go vet ./...
	golint -set_exit_status ./...

.PHONY: testdeps
testdeps:
	go get -d -v -t ./...
	GO111MODULE=off go get golang.org/x/lint/golint

.PHONY: clean
clean:
	go clean ./...
