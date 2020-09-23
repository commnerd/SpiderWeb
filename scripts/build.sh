#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

docker run -v ~/.go:/go -v ${DIR}/../cli:/workspace -w /workspace -e  CGO_ENABLED=0 -it --rm --name sw-SpiderWeb-cli-builder golang bash -c "go get -d && go build -tags netgo -a -o sweb"
