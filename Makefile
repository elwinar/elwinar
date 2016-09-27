.PHONY: dependencies
dependencies:
	gb vendor restore
	npm install
	bower install

.PHONY: build
build:
	node_modules/.bin/gulp
	go-bindata -nomemcopy -pkg main -o src/app/views.go src/views/
	rambler apply
	gb build -ldflags "-s -linkmode external -extldflags -static -w" -tags "docker"
	mv bin/app-docker build/app

.PHONY: dist
dist:
	goupx -q build/app
	docker build -t elwinar/elwinar .
	docker save elwinar/elwinar > build/elwinar.tar

.PHONY: clean
clean:
	rm -rf src/app/views pkg bin build

.PHONY: mrproper
mrproper:
	rm -rf node_modules bower_components vendor/src

.DEFAULT: build
