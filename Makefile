API_ONLY_PKGS=$(shell go list ./... 2> /dev/null | grep -v "/vendor/")

run:
	@go run main.go

test:
	@go test $(API_ONLY_PKGS)
