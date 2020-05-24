GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

fmt:
	gofmt -w $(GOFMT_FILES)

lint:
	golangci-lint run -v -c .golangci.yml ./...

test:
	go test ./... -timeout=2m -parallel=4

all: fmt lint tmp