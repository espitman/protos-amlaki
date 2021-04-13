#!/bin/bash 

protoc --go_out=plugins=grpc:. user/user.proto \
&& protoc-go-inject-tag -input=./user/user.pb.go \
&& git add . && git commit -m 'user update' && git push