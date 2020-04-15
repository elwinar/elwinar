default: build
all: dependencies build

.PHONY: dependencies
dependencies:
	npm install
	go get github.com/jteeuwen/go-bindata/...

.PHONY: build
build:
	./node_modules/.bin/parcel build -d build/public/ src/styles/elwinar.less
	go-bindata -nomemcopy -pkg main -o src/app/views.go src/views/
	go build -o build/app ./src/app

