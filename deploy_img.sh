#!/bin/bash

## make new image and run container
docker compose up --build -d

## save new image with different name
docker rmi backend-app:v1.0
docker tag backend-go-demo-gin-app:latest backend-app:v1.0

## export image as file
#docker save backend-app > backend-app-img.tar.gz