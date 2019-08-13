#!/bin/bash -e
path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go)
	go generate .
        goimports -w $path
        gofmt -w $path
        ;;
esac
go test -coverprofile /tmp/c.out ./...
#uncover /tmp/c.out
cp reference.html /var/www/html/demo/index.html
