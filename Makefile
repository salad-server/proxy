# Run without building (development)
serve:
	echo "serving proxy..."
	go run *.go

# Build
build:
	echo "building proxy..."
	go build

# Build with upx (compress output)
build-prod:
	echo "building proxy... (with upx)"
	go build
	upx proxy
