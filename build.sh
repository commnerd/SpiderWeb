#!/bin/bash

docker run -v ~/.go:/go -v ${PWD}:/go/src/github.com/commnerd/SpiderWeb -w /go/src/github.com/commnerd/SpiderWeb -e  CGO_ENABLED=0 -it --rm --name sw-SpiderWeb-builder golang bash -c "go get -d && go build -tags netgo -a -o sw-node"
