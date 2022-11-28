run: install
	$(GOPATH)/bin/api
install:
	@go install github.com/53jk1/go-api-boilerplate/cmd/api