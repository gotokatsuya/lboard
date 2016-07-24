API_ONLY_PKGS=$(shell go list ./... 2> /dev/null | grep -v "/vendor/")

run:
	@go run main.go

test:
	@go test $(API_ONLY_PKGS)

vet:
	@echo "go vet packages"
	@go tool vet -all -structtags -shadow $(shell ls -d /*/ | grep -v "vendor")
