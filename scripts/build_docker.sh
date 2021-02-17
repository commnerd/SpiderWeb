#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd ${DIR}/../worker_node

docker build --tag=commnerd/spiderweb_worker_node .
