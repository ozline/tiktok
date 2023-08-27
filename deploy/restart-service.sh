#!/bin/sh
source remove-all-containers.sh

SERVICE_TO_START=${1:-api} # get specific service

echo "Starting container for $SERVICE_TO_START..."
docker run -d --name "tiktok-$SERVICE_TO_START" \
-e service=$SERVICE_TO_START \
--net=host \
-v $DIR/config:/app/config \
"$IMAGE_NAME"