#!/usr/bin/env bash

python2 -m grpc_tools.protoc -I../../proto \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --python_out=. --grpc_python_out=. ../../proto/grpcbasic0.proto
