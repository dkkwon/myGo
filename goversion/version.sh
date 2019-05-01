#!/bin/bash

GIT_VER=`git describe --tags`
# GIT_REV='git rev-parse --short HEAD'
BUILDTIME=`date -I`

echo ${GIT_VER}
echo ${BUILDTIME}

go build -ldflags "-X main.version=${GIT_VER} -X main.buildTime=${BUILDTIME}"
