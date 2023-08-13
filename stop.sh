# Please do not modify this file, modify the .env file within this directory
# If you are unable to run this file then run
# chmod +x ./setdotenv

docker stop blogaggregator-database-1
docker stop blog_aggregator

docker rm blogaggregator-database-1
docker rm blog_aggregator

#removes images
if [ $(docker inspect blogaggregator-web) == [] ]; then
  echo "'blogaggregator-web' does not exist."
else
  docker rmi $(docker images 'blogaggregator-web' -a -q) -f
fi