#!/usr/bin/env make
# ---------------------------------------------------------------------------------------------------------------------

GO = go
PROTOC = protoc

# ---------------------------------------------------------------------------------------------------------------------

PBFLAGS ?=
PBFLAGS += -I "."
PBFLAGS += -I "/usr/local/include"
PBFLAGS += -I "$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway"
PBFLAGS += -I "$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis"

PBPLUGINS ?=
PBPLUGINS += --go_out=plugins=grpc:"$(GOPATH)/src"

PBPLUGINS_API ?=
PBPLUGINS_API += --grpc-gateway_out=logtostderr=true:"$(GOPATH)/src"
PBPLUGINS_API += --swagger_out=logtostderr=true,allow_merge=true,merge_file_name=OpenAPI:"./gen/api"
PBPLUGINS_API += $(PBPLUGINS)

# ---------------------------------------------------------------------------------------------------------------------

.PHONY: pb-build-rpc
pb-build-rpc:
	$(PROTOC) $(PBFLAGS) $(PBPLUGINS) --proto_path=./proto ./proto/*.proto

.PHONY: pb-build-api
pb-build-api:
	$(PROTOC) $(PBFLAGS) $(PBPLUGINS_API) --proto_path=./proto ./proto/*.proto

.PHONY: pb-build
pb-build: pb-build-api

# ---------------------------------------------------------------------------------------------------------------------

.PHONY: pb-deps
pb-deps:
	$(GO) get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GO) get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GO) get -u -v github.com/golang/protobuf/protoc-gen-go

# ---------------------------------------------------------------------------------------------------------------------
