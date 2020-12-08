#!/bin/bash

DIR=$(dirname $(dirname $(readlink -f $0)))
ENV=$1 || 'prod'

case $ENV in
    "prod") ENV="prod" ;;
    "local") ENV="local" ;;
    *) ENV="test" ;;
esac

sudo rm -fR $DIR/bin/*

docker run -v ${DIR}:/workspace -e -it --rm --name sw-SpiderWeb-ui-builder node bash -c "cd /workspace/frontend && yarn && yarn build"

if [[ -d $DIR/bin/dist ]]
then
    sudo rm -fR $DIR/bin/frontend
fi

sudo mv $DIR/frontend/dist/frontend $DIR/bin

docker run -v ~/.go:/go -v ${DIR}:/workspace -w /workspace/daemon -e  CGO_ENABLED=0 -it --rm --name sw-SpiderWeb-daemon-builder golang bash -c "go get -d && go build -tags api -a -o ../bin/swd"

sudo chmod +x $DIR/bin/swd

# if [[ "local" == $ENV ]]
# then
#     sudo mv $DIR/bin/swd $DIR/scripts/docker
#     sudo mv $DIR/bin/frontend $DIR/scripts/docker

#     cd $DIR/scripts/docker
#     docker build --tag commnerd/swd:latest .

#     sudo rm $DIR/scripts/docker/swd
#     sudo rm -fR $DIR/scripts/docker/frontend

#     docker run -d --name swd commnerd/swd
# fi

if [[ "local" == $ENV ]]
then
    cd $DIR/frontend && yarn && yarn build --watch --output-path="../bin/frontend" &
    P1=$!

    cd $DIR/bin
    sudo ./swd &
    P2=$!

    wait $P1 $p2
fi