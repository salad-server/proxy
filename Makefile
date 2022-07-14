# Run without building
serve:
	go run *.go

# Install the proxy
install:
	cp ext/config.example.json config.json
	go build; upx proxy
	nano config.json

# Build and compress
build:
	go build
	upx proxy
