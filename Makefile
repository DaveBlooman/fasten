default: install

install:
	go install
	fasten deploy

bindata:
	go-bindata -pkg command -o command/bindata.go libraries/...
