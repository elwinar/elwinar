build: assets app

assets:
	go get ./...
	npm install
	bower install

app:
	go-bindata -nomemcopy -debug -pkg main -o views.go views/
	go build -ldflags "-s -linkmode external -extldflags -static -w" -o elwinar
	gulp
	touch database.sqlite
	rambler apply --all

pkg: 
	go-bindata -nomemcopy -pkg main -o views.go views/
	go build -tags docker -ldflags "-s -linkmode external -extldflags -static -w" -o elwinar
	gulp
	touch database.sqlite
	rambler apply --all
	goupx -q elwinar
	docker build -t elwinar/elwinar .
	docker save elwinar/elwinar > elwinar.tar

clean:
	rm -f elwinar views.go public/*.js public/*.css database.sqlite
	rm -rf node_modules bower_components public/fonts

