#!/bin/bash

# Please do not modify this file, modify the .env file within this directory
# If you are unable to run this file then run
# chmod +x ./run.sh

OS="$(uname -s)"

# Builds the go code depending of OS
if [ $OS == "Darwin" ]; then
    echo "OSX detected"
    echo "GOOS=linux GOARCH=amd64 go build -o out"
    GOOS=linux GOARCH=amd64 go build -o out
else
    echo "Linux detected"
    echo "running go build -o out"
    go build -o out
fi

docker-compose rm -f

docker compose up -d
DOCKER_CONTAINER_NAME="blog_db"
until 
    docker exec $DOCKER_CONTAINER_NAME pg_isready;
do 
    sleep 3; 
done

