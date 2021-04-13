#!/bin/bash 

protoc --go_out=plugins=grpc:. neighborhood/neighborhood.proto \
&& protoc-go-inject-tag -input=./neighborhood/neighborhood.pb.go \
&& git add . && git commit -m 'neighborhood update' && git push