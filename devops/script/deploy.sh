#!/bin/bash

set -e

# Download and install Google Cloud SDK
curl -0 https://storage.googleapis.com/cloud-sdk-release/google-cloud-sdk-159.0.0-linux-x86_64.tar.gz | tar -zx -C ${HOME}
${HOME}/google-cloud-sdk/install.sh
source ${HOME}/google-cloud-sdk/path.bash.inc

# Install kubectl
gcloud --quiet components install kubectl

# Auth with service account
echo $GCLOUD_SERVICE_KEY | base64 --decode -i > ${HOME}/gcloud-service-key.json
gcloud auth activate-service-account --key-file ${HOME}/gcloud-service-key.json

gcloud --quiet config set project $PROJECT_NAME
gcloud --quiet config set container/cluster $CLUSTER_NAME
gcloud --quiet config set compute/zone ${CLOUDSDK_COMPUTE_ZONE}
gcloud --quiet container clusters get-credentials $CLUSTER_NAME

# Fix for gsutil (https://github.com/travis-ci/travis-ci/issues/7940)
export BOTO_CONFIG=/dev/null

# Assets
npm run bundle
gsutil -m rsync -r static/ gs://${PROJECT_NAME}/brynnse/${COMMIT}/static

gsutil -m setmeta \
  -h "Cache-Control:public, max-age=31536000" \
  -h "Content-Encoding: gzip" \
  gs://${PROJECT_NAME}/brynnse/${COMMIT}/static/**

# Build
WEB_IMAGE=eu.gcr.io/${PROJECT_NAME}/${DOCKER_IMAGE_NAME}/${K8S_DEPLOYMENT_NAME_WEB}
NGINX_IMAGE=eu.gcr.io/${PROJECT_NAME}/${DOCKER_IMAGE_NAME}/${K8S_DEPLOYMENT_NAME_NGINX}

# Build web image
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/main .
curl -o .build/ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
docker build \
  --build-arg STATIC=${STATIC}/$COMMIT/static \
  -t $WEB_IMAGE:$COMMIT \
  -t $WEB_IMAGE:latest \
  -f Dockerfile.scratch .

# Build nginx image
docker build \
  -t $NGINX_IMAGE:$COMMIT \
  -t $NGINX_IMAGE:latest \
  -f Dockerfile.nginx .

# Push images
gcloud docker -- push $WEB_IMAGE:$COMMIT
gcloud docker -- push $WEB_IMAGE:latest
gcloud docker -- push $NGINX_IMAGE:$COMMIT
gcloud docker -- push $NGINX_IMAGE:latest

# Update web
kubectl -n ${K8S_NAMESPACE} set image deployment/${K8S_DEPLOYMENT_NAME_WEB} ${K8S_DEPLOYMENT_NAME_WEB}=$WEB_IMAGE:$COMMIT

# Update nginx
kubectl -n ${K8S_NAMESPACE} set image deployment/${K8S_DEPLOYMENT_NAME_NGINX} ${K8S_DEPLOYMENT_NAME_NGINX}=$NGINX_IMAGE:$COMMIT
