.PHONY: all

all: clean build run

.PHONY: clean

clean:
	@rm -rf ./webview-react.app ./ui/build
	@echo "[✔️] Clean complete!"

.PHONY: build

build:
	@cd ./ui && npm install
	@cd ./ui && npm run build
	@mkdir -p ./webview-react.app/Contents/MacOS
	@go build -o ./webview-react.app/Contents/MacOS/webview-react
	@echo "[✔️] Build complete!"

.PHONY: run

run:
	@open ./webview-react.app
	@echo "[✔️] App is running!"
