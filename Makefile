linter:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

lint: linter
	gometalinter --enable=gofmt --vendor --deadline=60s ./...

test:
	go test -covermode=count -coverprofile=count.out -v ./...
