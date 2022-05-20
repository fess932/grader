.PHONY: server static grader gen

server: static
	go run ./cmd/server

grader:
	go run ./cmd/grader

#$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
#$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
gen:
	@protoc --version
	@protoc-gen-go --version
	@protoc-gen-go-grpc --version

	protoc --go_out=. --go-grpc_out=. api/proto/grader.proto

runner:
	docker build -f ./Dockerfile.runner -t runner .

run-go: runner
	docker run \
	-v $(shell pwd)/examples/go:/langs/go \
	-v $(shell pwd)/examples/go.yaml:/langs/tests.yaml \
	runner

run-python: runner
	docker run \
	-v $(shell pwd)/examples/python:/langs/python \
	-v $(shell pwd)/examples/python.yaml:/langs/tests.yaml \
	runner

static:
	npm run build:css