default: build
all: dependencies build

.PHONY: dependencies
dependencies:
	gb vendor restore
	npm install
	node_modules/.bin/bower install
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elwinar/rambler
	go get github.com/pwaller/goupx

.PHONY: build
build:
	node_modules/.bin/gulp
	rambler apply --all
	go-bindata -nomemcopy -pkg main -o src/app/views.go src/views/
	gb build -ldflags "-s -linkmode external -extldflags -static -w"
	goupx -q bin/app
	mv bin/app build/app

.PHONY: clean
clean:
	rm -rf src/app/views pkg bin build

.PHONY: mrproper
mrproper:
	rm -rf node_modules bower_components vendor/src

