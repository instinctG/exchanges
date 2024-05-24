test:
	go test -v ./...

server:
	go run ./...

.PHONY: test server