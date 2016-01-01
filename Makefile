dependencies: embeded *.go package.json bower.json
	go get ./...
	npm install
	bower install

assets: scripts/*.js styles/*.less
	gulp

database: migrations/*.sql
	touch database.sqlite
	rambler apply --all

embeded: 
	go-bindata -nomemcopy -debug -pkg main -o views.go views/

embeded-static:
	go-bindata -nomemcopy -pkg main -o views.go views/

binary: *.go
	go build -o elwinar

binary-static:
	go build -o elwinar -ldflags "-s -linkmode external -extldflags -static -w" -tags docker
	goupx -q elwinar

pkg: dependencies assets embeded-static binary-static
	docker build -t elwinar/elwinar .
	docker save elwinar/elwinar > elwinar.tar

clean:
	rm -f elwinar views.go public/*.js public/*.css database.sqlite elwinar.tar
	rm -rf node_modules bower_components public/fonts

