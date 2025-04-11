LOCAL_BIN:=$(CURDIR)/bin

.PHONY: deps
deps:
	$(info #Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
	GOBIN=$(LOCAL_BIN) go install github.com/mjibson/esc@v0.2.0

.PHONY: update
update: proto

# runs buf generate with installation of all deps
.PHONY: proto
proto:
	buf dep update
	buf dep prune
	buf lint
	#buf breaking --against ".git#subdir=."
	buf generate