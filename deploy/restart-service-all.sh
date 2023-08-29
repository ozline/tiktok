#!/bin/bash
source remove-all-containers.sh

echo "Pulling the latest image..."
docker pull "$IMAGE_NAME"

echo "Launch all services"
for service in api user chat follow interaction video; do
    echo "Starting container for $service..."
    docker run -d --name "tiktok-$service" \
    -e service=$service \
    --net=host \
    -v $DIR/config:/app/config \
    "$IMAGE_NAME"
done