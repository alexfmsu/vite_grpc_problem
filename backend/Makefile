generate=bash generate.sh

install:
	@go mod init backend 2>/dev/null || true
	@go get google.golang.org/genproto@latest
	@GOPROXY=https://goproxy.cn go get github.com/klauspost/compress@latest
	@go mod tidy

run:
	${generate}
	go run server.go -p 9002

clean:
	@rm go.mod so.sum 2>/dev/null || true
	@rm -rf proto/*.go proto/lotspb 2>/dev/null || true
