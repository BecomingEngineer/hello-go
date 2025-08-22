.PHONY: run test vet fmt check

run:
	go run .

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -s -w .

check: fmt vet test
