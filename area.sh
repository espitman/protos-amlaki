#!/bin/bash 

protoc --go_out=plugins=grpc:. area/area.proto \
&& protoc-go-inject-tag -input=./area/area.pb.go \
&& git add . && git commit -m 'area update' && git push