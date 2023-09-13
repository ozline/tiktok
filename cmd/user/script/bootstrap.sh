#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
CONFIG_PATH=$(dirname $(dirname $CURDIR))/config
REMOTE_CONFIG_PATH = /var/lib/etcd/conf/config.yaml
ENDPOINT = $1
RUNTIME_ROOT=${CURDIR}
#if [ "X$1" != "X" ]; then
#    RUNTIME_ROOT=$1
#else
#    RUNTIME_ROOT=${CURDIR}
#fi

export KITEX_RUNTIME_ROOT=$RUNTIME_ROOT
export KITEX_LOG_DIR="$RUNTIME_ROOT/log"

if [ ! -d "$KITEX_LOG_DIR/app" ]; then
    mkdir -p "$KITEX_LOG_DIR/app"
fi

if [ ! -d "$KITEX_LOG_DIR/rpc" ]; then
    mkdir -p "$KITEX_LOG_DIR/rpc"
fi

exec "$CURDIR/bin/user" -config $CONFIG_PATH -rc $REMOTE_CONFIG_PATH -e $ENDPOINT
