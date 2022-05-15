.PHONY: server static static-watch

server: static
	go run ./cmd/server

static:
	npm run build:css