debug-docker:
	go run main.go ./testdata/DL3000_Dockerfile

test:
	go test -cover -v $$(go list ./...)

