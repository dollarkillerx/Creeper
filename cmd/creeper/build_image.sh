#!/bin/sh

source_path=./cmd/creeper/
go_file=main.go
image_name=creeper
build_output=creeper
version=latest

CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -o $source_path/$build_output $source_path/$go_file

docker rmi -f $image_name:$version
docker build -f $source_path/Dockerfile -t $image_name:$version  .
rm $source_path/$build_output

