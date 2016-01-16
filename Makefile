target=release

build: 
	@echo '====== Build api ====='
	go build -ldflags "-s" -o $(target)/itpkg api/main.go
	mkdir -p $(target)/config
	cp -a api/config/development.toml $(target)/config
	cp -a api/templates .env $(target)/
	@echo '====== Build front ====='
	cd front-react && npm run build
	cp -a front-react/dist $(target)/public
	@echo '====== Build locales ====='
	cd tools && rake locales 
	mkdir -p $(target)/tmp
	cp tools/locales.sql $(target)/tmp


clean:
	-rm -r $(target)


format:
	for f in `find . -type f -iname '*.go'`; do gofmt -w $$f; done
