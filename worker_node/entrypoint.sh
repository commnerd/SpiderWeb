#!/bin/bash

set -eu

cp -fR /root/user_ssh /root/.ssh
chown -fR root:root /root/.ssh

#if [ ! -f /var/run/spiderweb/token ]
#then
	
   # kubeadm join --discovery-token-unsafe-skip-ca-verification spiderweb.com:6443

#fi

exec "$@"
