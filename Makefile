EXE=elwinar

build:
	go generate
	go build -ldflags "-s" -o $(EXE) . 
	goupx $(EXE)

