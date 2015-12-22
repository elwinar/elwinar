build: assets app

assets: *.go package.json bower.json
	go-bindata -nomemcopy -pkg main -o views.go views/
	go get ./...
	npm install
	bower install

app: *.go scripts/*.js styles/*.less
	go build -ldflags "-s" -o elwinar . 
	gulp
	touch database.sqlite
	rambler apply --all

pkg:
	goupx -q elwinar
	docker build -t elwinar .
	docker save elwinar > elwinar.tar

clean:
	rm -f elwinar views.go public/*.js public/*.css database.sqlite
	rm -rf node_modules bower_components public/fonts

