#!/bin/bash

set -e

# Assets
npm run bundle
gsutil -m cp -r -z js,css static gs://${PROJECT_NAME}/brynnse/${COMMIT}/static

gsutil -m setmeta \
  -h "Cache-Control:public, max-age=31536000" \
  gs://${PROJECT_NAME}/brynnse/${COMMIT}/static/**
