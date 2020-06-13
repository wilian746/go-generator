GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

fmt:
	gofmt -w $(GOFMT_FILES)

lint:
	golangci-lint run -v -c .golangci.yml ./...

test:
	go test ./... -timeout=2m -parallel=4

build:
	chmod +x ./deployments/scripts/setup_version.sh
	./deployments/scripts/setup_version.sh "no-changes"
	go build -o go-generator.tmp ./cmd/go-generator/main.go
	./deployments/scripts/setup_version.sh "rollback"

all: fmt lint test build