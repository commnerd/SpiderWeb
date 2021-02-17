#!/bin/bash

set -eu

cp -fR /root/user_ssh /root/.ssh
chown -fR root:root /root/.ssh

#if [ ! -f /var/run/spiderweb/token ]
#then
	
   # kubeadm join --discovery-token-unsafe-skip-ca-verification spiderweb.com:6443
   #kubeadm join 54.203.223.214:6443 --token mxg2d6.dz3j0jrp659qqe3z \
   # --discovery-token-ca-cert-hash sha256:e8203a6db3796219bdcbd4bf319393ef9c47af5a06a459e18ffe519e7a919daf

#fi



exec "$@"
