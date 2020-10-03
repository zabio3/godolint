.PHONY: check
check: fmt test

fmt:
	@go fmt ./...

test:
	@go test -cover -v ./...

docker:
	docker build --rm -t zabio3/godolint .
