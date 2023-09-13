#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
CONFIG_PATH=$(dirname $(dirname $CURDIR))/config
REMOTE_CONFIG_PATH = /var/lib/etcd/conf/config.yaml
ENDPOINT = $1

BinaryName=api
echo "$CURDIR/bin/${BinaryName}"

exec $CURDIR/bin/${BinaryName} -config $CONFIG_PATH -rc $REMOTE_CONFIG_PATH -e $ENDPOINT
