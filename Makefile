build:
	GOOS=js GOARCH=wasm go build -o interpreter.wasm
	cp -fr interpreter.wasm ./editor/static/wasm

start-dev: build
	cd ./editor && npm run dev
	
	