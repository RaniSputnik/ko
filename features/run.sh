#!/bin/bash

NETWORK_NAME=ko-tests-net

MYSQL_CONTAINER_NAME=ko-tests-mysql
MYSQL_ROOT_PASSWORD=example
MYSQL_DATABASE=ko

KO_API_CONTAINER_NAME=ko-tests-api

# Perform cleanup from previous run

echo 'Cleaning up old containers...'
docker stop $KO_API_CONTAINER_NAME
docker stop $MYSQL_CONTAINER_NAME
docker rm $KO_API_CONTAINER_NAME
docker rm $MYSQL_CONTAINER_NAME

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
# TODO find a better way to do this - possibly in application code?
echo 'Waiting for SQL to boot'
sleep 15

echo 'Running ko-app container...'
docker run --name $KO_API_CONTAINER_NAME -itd \
    --network=$NETWORK_NAME \
    -p 8080:8080 \
    -e KO_SQL_HOST=$MYSQL_CONTAINER_NAME \
    -e KO_SQL_PWD=$MYSQL_ROOT_PASSWORD \
    ko-app

# TODO run the tests

#docker stop $KO_API_CONTAINER_NAME
#docker stop $MYSQL_CONTAINER_NAME