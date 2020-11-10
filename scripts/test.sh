#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# docker run -v ~/.go:/go -v ${DIR}/..:/go/src/github.com/commnerd/SpiderWeb -v ${DIR}/..:/workspace -w /workspace/cli -e  CGO_ENABLED=0 -it --rm --name sw-SpiderWeb-cli-tester golang bash -c "go get -d && go test"
# docker run -v ~/.go:/go -v ${DIR}/..:/go/src/github.com/commnerd/SpiderWeb -v ${DIR}/..:/workspace -w /workspace/web -e  CGO_ENABLED=0 -it --rm --name sw-SpiderWeb-web-tester golang bash -c "go get -d && go test"

docker run -v ~/.go:/go -v ${DIR}/..:/go/src/github.com/commnerd/SpiderWeb -v ${DIR}/..:/workspace -it --rm --name sw-SpiderWeb-lib-tester golang bash -c "go get github.com/stretchr/testify/assert"
docker run -v ~/.go:/go -v ${DIR}/..:/go/src/github.com/commnerd/SpiderWeb -v ${DIR}/..:/workspace -it --rm --name sw-SpiderWeb-lib-tester golang bash -c "go get github.com/spf13/viper"
docker run -v ~/.go:/go -v ${DIR}/..:/go/src/github.com/commnerd/SpiderWeb -v ${DIR}/..:/workspace -it --rm --name sw-SpiderWeb-lib-tester golang bash -c "go get github.com/google/uuid"

libs=(
	"node"
	"message_bus"
	"config"
	"id"
	"services"
	"keys"
	"port"
	"api"
)

for l in ${libs[@]}; do
	docker run -v ~/.go:/go -v ${DIR}/..:/go/src/github.com/commnerd/SpiderWeb -v ${DIR}/..:/workspace -w /workspace/lib/$l -it --rm --name sw-SpiderWeb-lib-tester golang bash -c "go get -d && go test"
done
