#!/bin/bash

read -p "Enter version: " version
read -p "Server mode: " mode

date=`date +%Y-%m-%dT%H:%M:%S%Z`
commit=`git rev-parse --short HEAD`
root=`pwd`

GOOS=linux  GOARCH=amd64  go build -o ${root}/bin/mms -gcflags "-N -l" -ldflags "-X main.commit=${commit} -X main.version=${version} -X main.buildTime=${date} -X main.mode=${mode}"  ${root}/src/main/mms.go


