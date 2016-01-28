target=release

build: 
	@echo '====== Build locales ====='
	cd tools && rake locales 
	mkdir -p $(target)/tmp
	cp api/tmp/locales.txt $(target)/tmp
	@echo '====== Build api ====='
	go build -ldflags "-s" -o $(target)/itpkg api/main.go
	mkdir -p $(target)/config
	cp -a api/config/development.toml $(target)/config
	cp -a api/templates .env $(target)/
	@echo '====== Build front ====='
	cd front-react && rm -rf dist && npm run build
	cp -a front-react/dist $(target)/public


clean:
	-rm -r $(target)


format:
	for f in `find . -type f -iname '*.go'`; do gofmt -w $$f; done


locales:
	cd tools && rake locales
	cd api && go run main.go db s && go run main.go c c

reset:
	cd tools && rake locales
	cd api && go run main.go db d && go run main.go db n && go run main.go db m && go run main.go db s && go run main.go c c && go run main.go db c



