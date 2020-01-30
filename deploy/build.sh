#!/bin/sh

date=`date +%Y-%m-%dT%H:%M:%S%Z`
commit=`git rev-parse --short HEAD`
root=`pwd`

GOOS=linux  GOARCH=amd64  go build -v -o ${root}/bin/mms -gcflags "-N -l" -ldflags "-X main.commit=${commit} -X main.buildTime=${date}"  ${root}/main.go


