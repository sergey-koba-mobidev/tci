#!/bin/bash

# build linux
GOOS=linux GOARCH=amd64 go build -v .
mv tci builds/tci-Linux-x86_64

# build mac
GOOS=darwin GOARCH=amd64 go build -v .
mv tci builds/tci-Darwin-x86_64