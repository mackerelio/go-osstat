.PHONY: all
all: clean lint test

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint: testdeps
	golint -set_exit_status ./...

.PHONY: testdeps
testdeps:
	go install golang.org/x/lint/golint

.PHONY: clean
clean:
	go clean ./...
