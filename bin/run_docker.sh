#!/bin/bash

docker run \
	--privileged \
	-v /etc/hosts:/etc/hosts \
	-v /sys:/sys \
	-v ${HOME}/.ssh:/root/user_ssh \
	-d \
	commnerd/spiderweb_worker_node
