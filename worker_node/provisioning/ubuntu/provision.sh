#!/bin/bash

sudo su -

apt-get update && apt-get install -y \
    software-properties-common \
    apt-transport-https \
    ca-certificates \
    supervisor \
    docker.io \
    curl

# Set up the Docker daemon
cat <<EOF | tee /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

service docker restart

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list && \
    apt-get update && \
    apt-get install -y kubelet kubeadm kubectl && \
    apt-mark hold kubelet kubeadm kubectl

#kubeadm join spiderweb.com:6443 --token <token> \
#    --discovery-token-ca-cert-hash sha256:<SHA256>