.PHONY: prepare
prepare:
	./scripts/prepare.sh

.PHONY: test
test: prepare
	go test ./...

