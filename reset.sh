#!/bin/bash

# Please do not modify this file, modify the .env file within this directory
# If you are unable to run this file then run
# chmod +x ./reset.sh

source .env

docker stop $CONTAINER_NAME
docker stop $DB_NAME

docker rm $CONTAINER_NAME
docker rm $DB_NAME

#removes images
if docker image inspect $IMAGE_NAME >/dev/null 2>&1; then
  docker rmi $(docker images $IMAGE_NAME -a -q) -f
else
  echo "'blogaggregator-web' does not exist."
fi