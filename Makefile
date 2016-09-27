dependencies:
	gb vendor restore
	gb build github.com/jteeuwen/go-bindata/go-bindata
	gb build github.com/pwaller/goupx
	gb build github.com/elwinar/rambler
	npm install
	bower install

build:
	node_modules/.bin/gulp
	bin/go-bindata -nomemcopy -pkg main -o src/elwinar/views.go src/views/
	bin/rambler apply
	gb build -ldflags "-s -linkmode external -extldflags -static -w" -tags "docker"

dist:
	goupx -q bin/elwinar-docker
	docker build -t elwinar/elwinar .
	docker save elwinar/elwinar > elwinar.tar

clean:
	rm -f src/elwinar/views.go elwinar.tar
	rm -rf node_modules bower_components pkg bin vendor/src public

.DEFAULT: build
