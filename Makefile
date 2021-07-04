.PHONY: build
build:
	go build ./cmd/tgame

.PHONY: test
test:
	go test -v -race -cover