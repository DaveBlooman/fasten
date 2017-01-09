default: install

deploy:
	go install
	fasten deploy

install:
	go install
	fasten install

status:
	go install
	fasten status

bindata:
	go-bindata -pkg command -o command/bindata.go libraries/...
