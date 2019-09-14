.PHONY: check
check: fmt test

fmt:
	@go fmt ./...

test:
	@go test -cover -v ./...
