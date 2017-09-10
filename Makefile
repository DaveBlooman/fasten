default: build-local

bindata:
	go-bindata -pkg files -o files/bindata.go libraries/...

build:
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64"

build-local: bindata
	go build && go install
