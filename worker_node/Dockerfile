FROM ubuntu

RUN apt-get update && apt-get install -y \
    software-properties-common \
    apt-transport-https \
    ca-certificates \
    supervisor \
    curl

RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    echo "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable" && \
    apt-get update && \
    apt-get install -y docker.io && \
    apt-mark hold docker.io
    

RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list && \
    apt-get update && \
    apt-get install -y kubelet kubeadm kubectl && \
    apt-mark hold kubelet kubeadm kubectl

RUN mkdir -p /var/lib/kubelet

COPY entrypoint.sh /entrypoint.sh

COPY supervisor/* /etc/supervisor/conf.d/

COPY ./kubelet/config.yaml /var/lib/kubelet/config.yaml

ENTRYPOINT ["/entrypoint.sh"]

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf", "-n"]
