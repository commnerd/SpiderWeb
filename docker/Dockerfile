FROM alpine

# COPY sw-node /
# COPY .env.example /.env
RUN apk add supervisor openssh openssh-server docker

RUN mkdir -p /var/log/supervisor && \
    mkdir -p /var/run/spider_web

ADD configs/supervisor/supervisord.conf /etc/supervisor/
ADD configs/supervisor/conf.d/* /etc/supervisor/conf.d/

CMD ["supervisord", "-nc", "/etc/supervisor/supervisord.conf"]

