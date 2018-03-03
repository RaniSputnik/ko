#!/bin/bash

NETWORK_NAME=ko-tests-net

MYSQL_CONTAINER_NAME=ko-tests-mysql
MYSQL_ROOT_PASSWORD=example
MYSQL_DATABASE=ko

KO_API_CONTAINER_NAME=ko-tests-api

FEATURES_CONTAINER_NAME=ko-tests-features

# Perform cleanup from previous run

echo 'Cleaning up old API container...'
docker stop $KO_API_CONTAINER_NAME
docker rm $KO_API_CONTAINER_NAME

# Provision containers

docker network create --driver bridge $NETWORK_NAME

echo 'Running mysql container...'
docker run --name $MYSQL_CONTAINER_NAME -d \
    --network=$NETWORK_NAME \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
    -e MYSQL_DATABASE=$MYSQL_DATABASE \
    mysql

# Give the mysql container time to boot
# TODO currently we expose mysql port 3306 to host machine
# in order to run this test. We could instead try running
# a container within the network to ping mysql
echo 'Waiting for SQL to boot'
until nc -z localhost 3306; do sleep 1; echo "."; done

echo 'Building ko-app container...'
docker build -t ko-app .

echo 'Running ko-app container...'
docker run --name $KO_API_CONTAINER_NAME -itd \
    --network=$NETWORK_NAME \
    -p 8080:8080 \
    -e KO_SQL_HOST=$MYSQL_CONTAINER_NAME \
    -e KO_SQL_PWD=$MYSQL_ROOT_PASSWORD \
    ko-app

echo 'Building the ko-tests container...'
pushd ./features; docker build -t ko-tests .; popd

echo 'Running the feature tests...'
docker run --name $FEATURES_CONTAINER_NAME -it --rm \
    --network=$NETWORK_NAME \
    -e KO_SQL_HOST=$MYSQL_CONTAINER_NAME \
    -e KO_SQL_PWD=$MYSQL_ROOT_PASSWORD \
    -e KO_API_ENDPOINT=http://$KO_API_CONTAINER_NAME:8080/graphql \
    ko-tests

RESULT=$?

#docker stop $KO_API_CONTAINER_NAME
#docker stop $MYSQL_CONTAINER_NAME

# Ensure the correct exit code is returned
exit $RESULT