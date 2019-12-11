#!/bin/bash

read -p "Enter version: " version

date=`date +%Y-%m-%dT%H:%M:%S%Z`
commit=`git rev-parse --short HEAD`
root=`pwd`

go build -o ${root}/bin/mms -gcflags "-N -l" -ldflags "-X main.commit=${commit} -X main.version=${version} -X main.buildTime=${date}"  ${root}/src/main/mms.go


