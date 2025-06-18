#!/bin/bash

## store the image directly to minikube
eval $(minikube docker-env)
docker build -t backend-app:v1.0 .