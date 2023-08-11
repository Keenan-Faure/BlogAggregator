# Please do not modify this file, modify the .env file within this directory
# If you are unable to run this file then run
# chmod +x ./setdotenv

# sets the env variables
source .env

export APP_DB_URL=$DB_URL
export APP_DATABASE=$DATABASE
export APP_PORT=$PORT
export APP_WOO_STORE_NAME=$WOO_STORE_NAME
export APP_WOO_CONSUMER_KEY=$WOO_CONSUMER_KEY
export APP_WOO_CONSUMER_SECRET=$WOO_CONSUMER_SECRET
export STORE_NAME=$STORE_NAME
export APP_API_KEY=$API_KEY
export APP_API_PASSWORD=$API_PASSWORD
export APP_VERSION=$VERSION
export APP_T_DB=$T_DB
export APP_T_STORE_NAME=$T_STORE_NAME
export APP_T_API_KEY=$T_API_KEY
export APP_T_API_PASSWORD=$T_API_PASSWORD
export APP_T_VERSION=$T_VERSION
export APP_T_WOO_STORE_NAME=$T_WOO_STORE_NAME
export APP_T_WOO_CONSUMER_KEY=$T_WOO_CONSUMER_KEY
export APP_T_WOO_CONSUMER_SECRET=$T_WOO_CONSUMER_SECRET

printenv | sort | grep --color -E "APP_"

docker build . -t keenanfaure/blog_agg
docker run -d --env-file=.env -p 8080:8080 keenanfaure/blog_agg
