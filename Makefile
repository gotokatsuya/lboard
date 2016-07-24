
PRJDIR=github.com/gotokatsuya/lboard
PRJTARGET=lboard

BINDIR=misc/bin

API_ONLY_PKGS=$(shell go list ./... 2> /dev/null | grep -v "/vendor/")

build-js:
	cd asset && gulp build

build-go:
	go build -v -o $(BINDIR)/$(PRJTARGET) -i $(PRJDIR)
	
run:
	./$(BINDIR)/$(PRJTARGET)

test:
	go test $(API_ONLY_PKGS)
