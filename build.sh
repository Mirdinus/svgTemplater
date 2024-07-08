#!/bin/bash

OUT_DIR=build

rm -rf $OUT_DIR
mkdir $OUT_DIR

# LINUX
GOOS=linux GOARCH=amd64 go build -o $OUT_DIR/cfChecker-amd64
GOOS=linux GOARCH=arm64 go build -o $OUT_DIR/cfChecker-arm

# WIN
GOOS=windows GOARCH=amd64 go build -o $OUT_DIR/cfChecker-windows.exe

# OSX
GOOD=darwin GOARCH=amd64 go build -o $OUT_DIR/cfChecker-darwin
GOOD=darwin GOARCH=arm64 go build -o $OUT_DIR/cfChecker-darwin-arm