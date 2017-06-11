#!/bin/bash

set -e

IMAGE=eu.gcr.io/${PROJECT_NAME}/${DOCKER_IMAGE_NAME}/${K8S_DEPLOYMENT_NAME}
VERSION_IMAGE=$IMAGE:$COMMIT
LATEST_IMAGE=$IMAGE:latest

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/main .
curl -o .build/ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
docker build --build-arg CACHEBUST=$COMMIT \
  -t $VERSION_IMAGE \
  -f Dockerfile.web .

echo $GCLOUD_SERVICE_KEY | base64 --decode -i > ${HOME}/gcloud-service-key.json
gcloud auth activate-service-account --key-file ${HOME}/gcloud-service-key.json

gcloud --quiet config set project $PROJECT_NAME
gcloud --quiet config set container/cluster $CLUSTER_NAME
gcloud --quiet config set compute/zone ${CLOUDSDK_COMPUTE_ZONE}
gcloud --quiet container clusters get-credentials $CLUSTER_NAME

gcloud docker -- push $VERSION_IMAGE

yes | gcloud beta container images add-tag $VERSION_IMAGE $LATEST_IMAGE

kubectl config view
kubectl config current-context

kubectl -n ${K8S_NAMESPACE} set image deployment/${K8S_DEPLOYMENT_NAME} ${K8S_DEPLOYMENT_NAME}=$VERSION_IMAGE
