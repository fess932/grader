.PHONY: server static grader gen

server: static
	go run ./cmd/server

grader:
	go run ./cmd/grader

gen:
	@protoc --version
	@protoc-gen-go --version
	@protoc-gen-go-grpc --version

	protoc --go_out=. --go-grpc_out=. api/proto/grader.proto

runner:
	docker build -f ./Dockerfile.runner -t runner .

static:
	npm run build:css