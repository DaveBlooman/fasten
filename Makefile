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

stop:
	go install
	fasten stop

bindata:
	go-bindata -pkg files -o files/bindata.go libraries/...
