.PHONY: all
all: clean test

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean ./...
