#!/bin/bash
source common.sh

containers_to_stop=$(docker ps -aq --filter "ancestor=$IMAGE_NAME")

for container_id in $containers_to_stop; do
    container_status=$(docker inspect -f '{{.State.Status}}' "$container_id")
    if [ "$container_status" == "running" ]; then
        echo "Stopping container $container_id..."
        docker stop "$container_id"
    elif [ "$container_status" == "paused" ]; then
        echo "Unpausing and then stopping container $container_id..."
        docker unpause "$container_id"
        docker stop "$container_id"
    fi
    echo "Remove container $container_id..."
    docker rm "$container_id"
done