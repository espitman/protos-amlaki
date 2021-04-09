#!/bin/bash 

protoc --go_out=plugins=grpc:. house/house.proto \
&& protoc-go-inject-tag -input=./house/house.pb.go \
&& git add . && git commit -m 'house update' && git push