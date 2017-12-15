default: install

install:
	@go install .

uninstall:
	@rm $(GOPATH)/bin/csvConcat

test:
	@go test . $(OPTS)

sure:
	@go test -race .
	@go fmt .
	@go vet .
	@golint .
	@go install .

.PHONY: test
