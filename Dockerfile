FROM alpine

ENV NODE_ROLE=node

COPY sw-node /
COPY .env.example /.env
RUN apk add supervisor openssh openssh-server docker

RUN mkdir -p /var/log/supervisor
ADD configs/supervisor/supervisord.conf /etc/supervisor/
ADD configs/supervisor/conf.d/* /etc/supervisor/conf.d/

EXPOSE 80 22

CMD ["supervisord", "-nc", "/etc/supervisor/supervisord.conf"]

