#!/bin/bash

docker run -v ~/.go:/go -v ${PWD}:/workspace -w /workspace -e  CGO_ENABLED=0 -it --rm --name sw-SpiderWeb-builder golang bash -c "go get -d && go build -tags netgo -a -o sw-node"
