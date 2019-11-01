#ARG ARCH="amd64"
#ARG OS="linux"
FROM   quay.io/prometheus/busybox:latest
LABEL  maintainer="The Authors <hpy253215039@163.com>"

#ARG ARCH="amd64"
#ARG OS="linux"
COPY ./RocketmqExporter /bin/RocketmqExporter

USER        nobody
EXPOSE      9104
ENTRYPOINT  [ "/bin/RocketmqExporter" ]
