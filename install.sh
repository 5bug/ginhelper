#!/bin/bash

go install github.com/go-bindata/go-bindata@latest

echo "go-bindata pkg/*/**"
go-bindata pkg/*/**

echo "go build to /usr/local/bin/"
go build -o /usr/local/bin/ginhelper *.go
echo "success"
