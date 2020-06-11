#!/usr/bin/env make
# ---------------------------------------------------------------------------------------------------------------------

GO = go
PROTO = protoc

# ---------------------------------------------------------------------------------------------------------------------

PBFLAGS ?=
PBFLAGS += -I .
PBFLAGS += -I /usr/local/include
PBFLAGS += -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway
PBFLAGS += -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

PBPLUGINS ?=
PBPLUGINS += --grpc-gateway_out=logtostderr=true:$(GOPATH)/src
PBPLUGINS += --swagger_out=allow_merge=true,merge_file_name=OpenAPI:./gen/v1/api
PBPLUGINS += --go_out=plugins=grpc:$(GOPATH)/src

# ---------------------------------------------------------------------------------------------------------------------

.PHONY: go-proto-build
go-proto:
	$(PROTO) $(PBFLAGS) $(PBPLUGINS) --proto_path=./proto ./proto/**/*.proto

# ---------------------------------------------------------------------------------------------------------------------

.PHONY: go-proto-setup
go-proto-setup:
	$(GO) get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GO) get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GO) get -u -v github.com/golang/protobuf/protoc-gen-go

# ---------------------------------------------------------------------------------------------------------------------
