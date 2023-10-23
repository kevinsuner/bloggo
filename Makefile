BINARY_NAME := bloggo

all: clean build compress

build:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}-windows.exe .
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}-linux .
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}-darwin .

compress:
	upx --all-methods ./bin/${BINARY_NAME}-windows.exe
	upx --all-methods ./bin/${BINARY_NAME}-linux
	upx --all-methods ./bin/${BINARY_NAME}-darwin

clean:
	go clean
	rm ./bin/${BINARY_NAME}-windows.exe ./bin/${BINARY_NAME}-linux ./bin/${BINARY_NAME}-darwin