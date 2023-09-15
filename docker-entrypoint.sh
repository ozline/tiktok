#!/bin/sh

CONFIGPATH="./config/config.yaml"

function read_key() {
    local key="$2"
    local flag=0
    while read -r LINE; do
        if [[ $flag == 0 ]]; then
            if [[ "$LINE" == *"$key:"* ]]; then
                if [[ "$LINE" == *" "* ]]; then
                    echo "$LINE" | awk -F " " '{print $2}'
                    return
                else
                    continue
                fi
            fi
        fi
    done < "$1"
}


# JAEGER
export JAEGER_DISABLED=false
export JAEGER_SAMPLER_TYPE="const"
export JAEGER_SAMPLER_PARAM=1
export JAEGER_REPORTER_LOG_SPANS=true

export JAEGER_AGENT_HOST=$(read_key $CONFIGPATH "jaeger-host")
export JAEGER_AGENT_PORT=$(read_key $CONFIGPATH "jaeger-port")

export ETCD_ADDR=$(read_key $CONFIGPATH "etcd-addr")

sh ./output/${service}/bootstrap.sh