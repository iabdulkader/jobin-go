package_manager = pnpm

all: build

templ: tailwind
	templ generate -path ./views
	templ generate -path ./components

tailwind: init
	$(package_manager) run build:css

tailwind-watch:
	$(package_manager) run watch:css

run:
	templ generate --watch --cmd="make dev"

dev: 
	go run cmd/server/main.go

build: templ
	go build -o jobin cmd/server/main.go

init:
	$(package_manager) install
	cp ./node_modules/preline/dist/preline.js ./static/js/preline.js


