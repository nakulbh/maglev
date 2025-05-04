.PHONY: build clean coverage test run

run: build
	./bin/maglev -api-keys=test

build:
	go build -o bin/maglev ./cmd/api

clean:
	go clean
	rm -f maglev
	rm -f coverage.out

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test:
	go test ./...
