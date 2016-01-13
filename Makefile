target=release

build:
	go build -ldflags "-s" -o $(target)/itpkg api/main.go
	mkdir -p $(target)/config
	cp -a api/config/development.toml $(target)/config
	cp -a api/templates $(target)/
	mkdir -p $(target)/tmp
	cd front && ember build --environment="production"


clean:
	-rm -r $(target)


format:
	for f in `find . -type f -iname '*.go'`; do gofmt -w $$f; done





