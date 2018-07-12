#!/bin/bash

rm -rf distribution

mkdir distribution
mkdir distribution/windows
mkdir distribution/linux
mkdir distribution/mac

GOOS=windows GOARCH=386 go build -o distribution/windows/probe.exe -v .
GOOS=windows GOARCH=amd64 go build -o distribution/windows/probe-x64.exe -v .
GOOS=linux GOARCH=arm go build -o distribution/linux/probe -v .
GOOS=linux GOARCH=amd64 go build -o distribution/linux/probe-x64 -v .
GOOS=darwin GOARCH=amd64 go build -o distribution/mac/probe -v .
