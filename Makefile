#!/usr/bin/env make

compile:
	protoc -I . -I /usr/local/include --proto_path=. --go_out=plugins=grpc:./gen ./proto/*.proto
	#protoc -I . -I /usr/local/include --proto_path=. --go_out=$(GOPATH)/src ./proto/*.proto
