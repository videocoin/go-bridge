PROTOC_BUILD_COMMAND=docker run \
    -v $(shell pwd):/go_workspace/src/github.com/videocoin/go-payment-manager \
    -v $(shell pwd)/proto_gen.mk:/app/proto_gen.mk \
    -w /go_workspace/src/github.com/videocoin/go-payment-manager \
	gcr.io/videocoin-network/protobuf-build-env

.PHONY: protoc
protoc:
	${PROTOC_BUILD_COMMAND} --target protoc
