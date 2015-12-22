build:
	go-bindata -nomemcopy -pkg main -o views.go views/
	go build -ldflags "-s" -o elwinar . 
	goupx elwinar
	gulp

npm: package.json
	npm install

bower: bower.json
	bower install

database:
	touch elwinar.sqlite
	rambler apply --all

clean:
	rm -f elwinar views.go public/*.js public/*.css elwinar.sqlite
	rm -rf node_modules bower_components public/fonts

