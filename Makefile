.PHONY: server static static-watch

server: static
	go run ./cmd/server

gen:
	@protoc --version
	@protoc-gen-go --version
	@protoc-gen-go-grpc --version

	protoc --go_out=. --go-grpc_out=. api/proto/grader.proto

static:
	npm run build:css