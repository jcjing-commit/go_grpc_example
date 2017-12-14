#!/usr/bin/env bash

export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH/bin
rm -rf bin
rm -rf pkg
cd ./src
rm -rf github.com
rm -rf golang.org
rm -rf google.golang.org
cd ..
pwd
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
cd ./src
pwd
go get -u google.golang.org/grpc

protoc --go_out=plugins=grpc:. ./proto/calc.proto
echo "OK"






