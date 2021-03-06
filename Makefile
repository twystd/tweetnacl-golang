all: test      \
	 benchmark \
     coverage

format: 
	go fmt ./...

build: format
	go build ./...

# build-all: format
#	env GOOS=linux   GOARCH=amd64  go build ./...
#	env GOOS=darwin  GOARCH=amd64  go build ./...
#	env GOOS=windows GOARCH=amd64  go build ./...

test: build
	go test ./...

vet: build
	go vet ./...

lint: build
	golint ./...

benchmark: build
	go test -bench ./...

coverage: build
	go test -cover ./...

debug: build
	gofmt -w=true *.go
	go test -run TestCryptoHashBlocks

