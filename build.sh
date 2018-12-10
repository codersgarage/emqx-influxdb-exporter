#!/bin/bash

export CONSUL_URL="localhost:8500"
export CONSUL_PATH="golang-restful-boilerplate"

export GOARCH="amd64"
export GOOS="darwin"
#export GOOS="linux"
export CGO_ENABLED=0

cmd=$1

package="github.com/codersgarage/emqx-influxdb-exporter"
binary="golang-restful-boilerplate"

if [ "$cmd" = "run" ]; then
    echo "Executing run command"
    ./${binary} serve
    exit;
fi

if [ "$cmd" = "build" ]; then
    echo "Executing build command"
    glide install
    go build -v -o ${binary}
    exit;
fi

if [ "$cmd" = "spew" ]; then
    echo "Executing spew command"
    glide install
    go build -v -o ${binary}
    ./${binary} serve
    exit;
fi

if [ "$cmd" = "m_create" ]; then
    echo "Executing migration create command"
    ./${binary} migration create
    exit;
fi

if [ "$cmd" = "m_drop" ]; then
    echo "Executing migration drop command"
    ./${binary} migration create
    exit;
fi

if [ "$cmd" = "m_auto" ]; then
    echo "Executing migration auto command"
    ./${binary} migration auto
    exit;
fi

echo "No command specified"
