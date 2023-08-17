
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
CONFIG_PATH=$(dirname $(dirname $(dirname $CURDIR)))/config
BinaryName=api
echo "$CURDIR/bin/${BinaryName}"

export JAEGER_DISABLED=false
export JAEGER_SAMPLER_TYPE="const"
export JAEGER_SAMPLER_PARAM=1
export JAEGER_REPORTER_LOG_SPANS=true
export JAEGER_AGENT_HOST="127.0.0.1"
export JAEGER_AGENT_PORT=6831

exec $CURDIR/bin/${BinaryName} -config $CONFIG_PATH