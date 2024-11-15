.PHONY: coverage
coverage:
    go test -v -covermode=count -coverprofile=coverage.out $(go list ./... | grep -v cmd/scratch)
    go tool cover -func=coverage.out


![Coverage](https://github.com/abhinandkakkadi/golang-test/badge.svg)
