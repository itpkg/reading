target=release

build:
	#npm run build
	go build -ldflags "-s" -o $(target)/itpkg api/main.go
	mkdir -p $(target)/config
	cp -a api/config/development.toml $(target)/config
	cp -a templates $(target)
	mkdir -p $(target)/tmp


clean:
	-rm -r $(target)


format:
	for f in `find . -type f -iname '*.go'`; do gofmt -w $$f; done





