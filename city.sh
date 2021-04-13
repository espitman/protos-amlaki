#!/bin/bash 

protoc --go_out=plugins=grpc:. city/city.proto \
&& protoc-go-inject-tag -input=./city/city.pb.go \
&& git add . && git commit -m 'city update' && git push