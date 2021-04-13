#!/bin/bash 

protoc --go_out=plugins=grpc:. province/province.proto \
&& protoc-go-inject-tag -input=./province/province.pb.go \
&& git add . && git commit -m 'province update' && git push