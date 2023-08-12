# Please do not modify this file, modify the .env file within this directory
# If you are unable to run this file then run
# chmod +x ./setdotenv

echo "running go build -o out"
go build -o out

docker-compose rm -f

docker compose up -d
