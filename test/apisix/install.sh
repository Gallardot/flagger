#!/usr/bin/env bash

set -o errexit

APISIX_CHART_VERSION="0.11.3" # apisix 2.15.1
REPO_ROOT=$(git rev-parse --show-toplevel)

mkdir -p ${REPO_ROOT}/bin

echo '>>> Creating apisix namespace'
kubectl create ns apisix

echo '>>> Installing APISIX'
helm repo add apisix https://charts.apiseven.com

helm upgrade -i apisix apisix/apisix --version=${APISIX_CHART_VERSION} \
--namespace apisix \
--set serviceMonitor.enabled=true \
--set apisix.podAnnotations."prometheus\.io/scrape"=true \
--set apisix.podAnnotations."prometheus\.io/port"=9091 \
--set apisix.podAnnotations."prometheus\.io/path"=/apisix/prometheus/metrics \
--set dashboard.enabled=true \
--set ingress-controller.enabled=true \
--set ingress-controller.config.apisix.serviceNamespace=apisix

kubectl -n apisix rollout status deployment/apisix
kubectl -n apisix get all

echo '>>> Installing Flagger'
helm upgrade -i flagger ${REPO_ROOT}/charts/flagger \
--set crd.create=false \
--namespace apisix \
--set prometheus.install=true \
--set meshProvider=apisix \
--set image.repository=test\/flagger \
--set image.tag=latest \

kubectl -n apisix get all
