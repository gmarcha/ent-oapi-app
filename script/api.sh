#! /bin/bash

CONTAINER_NAME=api

docker build -t $CONTAINER_NAME .
docker run -dp 8080:8080 $CONTAINER_NAME
