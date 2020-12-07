#!/bin/bash

DIR=$(dirname $(dirname $(readlink -f $0)))

docker run -v ~/.go:/go -v ${DIR}:/workspace -w /workspace/daemon -e  CGO_ENABLED=0 -it --rm --name sw-SpiderWeb-cli-builder golang bash -c "go get -d && go build -tags api -a -o ../bin/swd"
