#!/bin/bash

# Please do not modify this file, modify the .env file within this directory
# If you are unable to run this file then run
# chmod +x ./stop.sh

source .env

docker stop $CONTAINER_NAME
docker stop $DB_NAME

docker rm $CONTAINER_NAME
docker rm $DB_NAME