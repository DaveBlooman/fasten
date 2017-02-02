default: install

deploy:
	go install
	fasten deploy

install:
	go install
	fasten install

init:
	go install
	fasten init

status:
	go install
	fasten status

bindata:
	go-bindata -pkg files -o bindata/bindata.go libraries/...
