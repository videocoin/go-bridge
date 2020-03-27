REGISTRY=registry.dev.videocoin.net/bridge
VERSION ?= dev

.PHONY: generate
generate:
	go generate ./...

.PHONY: build
build:
	go build -o ./build/bridge ./cmd/

.PHONY: build-vendor
build-vendor:
	go build -o ./build/bridge -mod=vendor ./cmd/

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
	modvendor -copy="**/*.c **/*.h" -v

.PHONY: images
images:
	docker build -t ${REGISTRY}/bridge/bridge:$(VERSION) -f _assets/Dockerfile .

.PHONY: push
push:
	docker push ${REGISTRY}/bridge/bridge:$(VERSION)


.PHONY: test
test:
	go test ./...
