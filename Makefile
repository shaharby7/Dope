.PHONY: prepare
prepare:
	./hack/prepare.sh

.PHONY: test
test: prepare
	go test ./...

