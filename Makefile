build:
	GO111MODULE=on
	GOOS=js GOARCH=wasm go get
	GOOS=js GOARCH=wasm go build -o lib.wasm main.go
build-cli:
	GO111MODULE=on
	GOOS=linux GOARCH=amd64 go build -o oak cmd/oak-cli/main.go
		sudo cp oak /usr/local/bin
run:
	$(GOCMD) build 
