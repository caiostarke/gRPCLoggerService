#!/bin/bash

echo "protoc executed"
protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logger_pb/logger.proto