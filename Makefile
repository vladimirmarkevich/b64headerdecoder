all: test vet fmt build

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

build:
	go build -o bin/b64hd ./cmd/b64hd

run:
	./bin/b64hd ./test/sample-mails.txt
