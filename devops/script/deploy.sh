#!/bin/bash

set -e

WEB_IMAGE=eu.gcr.io/${PROJECT_NAME}/${DOCKER_IMAGE_NAME}/${K8S_DEPLOYMENT_NAME_WEB}
NGINX_IMAGE=eu.gcr.io/${PROJECT_NAME}/${DOCKER_IMAGE_NAME}/${K8S_DEPLOYMENT_NAME_NGINX}

# Build web
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/main .
curl -o .build/ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
docker build --build-arg CACHEBUST=$COMMIT \
  -t $WEB_IMAGE:$COMMIT \
  -t $WEB_IMAGE:latest \
  -f Dockerfile.web .

# Build nginx
docker build \
  -t $NGINX_IMAGE:$COMMIT \
  -t $NGINX_IMAGE:latest \
  -f Dockerfile.nginx .

# Push to Google Container Registry
gcloud docker -- push $WEB_IMAGE:$COMMIT
gcloud docker -- push $WEB_IMAGE:latest
gcloud docker -- push $NGINX_IMAGE:$COMMIT
gcloud docker -- push $NGINX_IMAGE:latest

# Update web
kubectl -n ${K8S_NAMESPACE} set image deployment/${K8S_DEPLOYMENT_NAME_WEB} ${K8S_DEPLOYMENT_NAME_WEB}=$WEB_IMAGE:$COMMIT

# Update nginx
kubectl -n ${K8S_NAMESPACE} set image deployment/${K8S_DEPLOYMENT_NAME_NGINX} ${K8S_DEPLOYMENT_NAME_NGINX}=$NGINX_IMAGE:$COMMIT
