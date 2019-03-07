

tests:
	go test -race  ./... -coverprofile=coverage.txt -covermode=atomic

vet:
	go vet $(go list ./... | grep -v /vendor/)


.PHONY: tests vet

