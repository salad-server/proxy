serve:
	go run *.go

install:
	cp config.example.json config.json
	go build; upx proxy
	nano config.json

build:
	go build
	upx proxy
