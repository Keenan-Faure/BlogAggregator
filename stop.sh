# Please do not modify this file, modify the .env file within this directory
# If you are unable to run this file then run
# chmod +x ./setdotenv

IMAGE_NAME=blogaggregator-web

docker stop blogaggregator-database-1
docker stop blog_aggregator

docker rm blogaggregator-database-1
docker rm blog_aggregator

#removes images
if docker image inspect $IMAGE_NAME >/dev/null 2>&1; then
  docker rmi $(docker images $IMAGE_NAME -a -q) -f
else
  echo "'blogaggregator-web' does not exist."
fi