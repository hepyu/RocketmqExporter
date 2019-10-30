ARG ARCH="amd64"
ARG OS="linux"
FROM   quay.io/prometheus/busybox:latest
LABEL  maintainer="The Authors <hpy253215039@163.com>"

ARG ARCH="amd64"
ARG OS="linux"
COPY ./rocketmq_exporter /bin/rocketmq_exporter

COPY ./env.default.config.backup ~/env.default.config.backup

RUN cat ~/env.default.config.backup >> ~/.bashrc

USER        nobody
EXPOSE      9104
ENTRYPOINT  [ "/bin/rocketmq_exporter" ]
