REGISTRY=registry.videocoin.net/bridge
VERSION ?= dev
TAGS ?= prometheus

.PHONY: generate
generate:
	go generate ./...

.PHONY: build
build:
	go build -tags=$(TAGS) -o ./build/bridge ./cmd/

.PHONY: build-vendor
build-vendor:
	go build -tags=$(TAGS) -o ./build/bridge -mod=vendor ./cmd/

.PHONY: tokenfct
tokenfct:
	go build -o ./build/tokenfct ./cmd/tools/tokenfct

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
	modvendor -copy="**/*.c **/*.h" -v

.PHONY: abi
abi:
	./_assets/solc.sh ./build/ $(shell pwd)/solidity/ $(shell pwd)/solidity/contracts/bridge/NativeBridge.sol
	./_assets/solc.sh ./build/ $(shell pwd)/solidity/ $(shell pwd)/solidity/contracts/bridge/NativeProxy.sol
	./_assets/solc.sh ./build/ $(shell pwd)/solidity/ $(shell pwd)/solidity/contracts/bridge/RemoteBridge.sol

.PHONY: images
images:
	docker build -t ${REGISTRY}/bridge:$(VERSION) -f _assets/Dockerfile .

.PHONY: push
push:
	docker push ${REGISTRY}/bridge:$(VERSION)


.PHONY: test
test:
	go test ./...
