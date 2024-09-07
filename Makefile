.PHONY: prepare
prepare:
	./hack/prepare-example.sh

.PHONY: test
test: prepare
	go test ./...

