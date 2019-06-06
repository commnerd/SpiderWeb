#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

docker run -v ~/.go:/go -v ${DIR}/..:/go/src/github.com/commnerd/SpiderWeb -v ${DIR}/..:/workspace -w /workspace/cli -e  CGO_ENABLED=0 -it --rm --name sw-SpiderWeb-cli-tester golang bash -c "go get -d && go test"
docker run -v ~/.go:/go -v ${DIR}/..:/go/src/github.com/commnerd/SpiderWeb -v ${DIR}/..:/workspace -w /workspace/web -e  CGO_ENABLED=0 -it --rm --name sw-SpiderWeb-web-tester golang bash -c "go get -d && go test"
