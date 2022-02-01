build:
	GOARCH=wasm GOOS=js go build -o docs/web/app.wasm
	go build -o docs/gen

run: build
	cd docs && ./gen