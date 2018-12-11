FROM golang:alpine

RUN apk add --no-cache --update git
ENV GOPATH=/go
COPY emqx-influxdb-exporter /usr/local/bin/

ENTRYPOINT ["emqx-influxdb-exporter"]
