#!/bin/sh

source_path=./cmd/rlog/
go_file=main.go
image_name=creeper_rlog
build_output=creeper_rlog
version=latest

CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -o $source_path/$build_output $source_path/$go_file

docker rmi -f $image_name:$version
docker build -f $source_path/Dockerfile -t $image_name:$version  .
rm $source_path/$build_output

