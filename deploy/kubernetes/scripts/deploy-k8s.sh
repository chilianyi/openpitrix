#!/bin/bash

DEFAULT_NAMESPACE=openpitrix-system
DEFAULT_VERSION=$(curl -L -s https://api.github.com/repos/openpitrix/openpitrix/releases/latest | grep tag_name | sed "s/ *\"tag_name\": *\"\(.*\)\",*/\1/")

NAMESPACE=${DEFAULT_NAMESPACE}
VERSION=${DEFAULT_VERSION}
METADATA=0
DBCTRL=0
BASE=0

usage() {
  echo "Usage:"
  echo "  deploy-k8s.sh [-n NAMESPACE] [-v VERSION] COMMAND"
  echo "Description:"
  echo "    -n NAMESPACE: the namespace of kubernetes."
  echo "    -v VERSION  : the version to be deployed."
  echo "    -b          : base model will be applied."
  echo "    -m          : metadata will be applied."
  echo "    -d          : dbctrl will be applied."
  exit -1
}

while getopts n:v:m:hbdm option
do
  case "${option}"
  in
  n) NAMESPACE=${OPTARG};;
  v) VERSION=${OPTARG};;
  d) DBCTRL=1;;
  m) METADATA=1;;
  b) BASE=1;;
  h) usage ;;
  *) usage ;;
  esac
done

if [ "${METADATA}" == "0" ] && [ "${DBCTRL}" == "0" ] && [ "${BASE}" == "0" ];then
  usage
fi

if [ "${VERSION}" == "dev" ];then
  IMAGE="openpitrix/openpitrix-dev:latest"
  METADATA_IMAGE="openpitrix/openpitrix-dev:metadata"
  FLYWAY_IMAGE="openpitrix/openpitrix-dev:flyway"
elif [ "$VERSION" == "latest" ];then
  IMAGE="openpitrix/openpitrix:latest"
  METADATA_IMAGE="openpitrix/openpitrix:metadata"
  FLYWAY_IMAGE="openpitrix/openpitrix:flyway"
else
  IMAGE="openpitrix/openpitrix:$VERSION"
  METADATA_IMAGE="openpitrix/openpitrix:metadata-$VERSION"
  FLYWAY_IMAGE="openpitrix/openpitrix:flyway-$VERSION"
fi

echo "Deploying k8s resource..."
# Back to the root of the project
cd $(dirname $0)
cd ../..

kubectl create namespace ${NAMESPACE}
kubectl create secret generic mysql-pass --from-file=./kubernetes/password.txt -n ${NAMESPACE}


if [ "${DBCTRL}" == "1" ];then
  for FILE in `ls ./kubernetes/ctrl`;do
    sed -e "s!\${NAMESPACE}!${NAMESPACE}!g" -e "s!\${IMAGE}!${IMAGE}!g" -e "s!\${METADATA_IMAGE}!${METADATA_IMAGE}!g" -e "s!\${FLYWAY_IMAGE}!${FLYWAY_IMAGE}!g" ./kubernetes/ctrl/${FILE} | kubectl delete -f - --ignore-not-found=true
    sed -e "s!\${NAMESPACE}!${NAMESPACE}!g" -e "s!\${IMAGE}!${IMAGE}!g" -e "s!\${METADATA_IMAGE}!${METADATA_IMAGE}!g" -e "s!\${FLYWAY_IMAGE}!${FLYWAY_IMAGE}!g" ./kubernetes/ctrl/${FILE} | kubectl apply -f -
  done
fi
if [ "${BASE}" == "1" ];then
  for FILE in `ls ./kubernetes/db/ | grep -v "\-job.yaml$"`;do
    sed -e "s!\${NAMESPACE}!${NAMESPACE}!g" -e "s!\${IMAGE}!${IMAGE}!g" -e "s!\${METADATA_IMAGE}!${METADATA_IMAGE}!g" -e "s!\${FLYWAY_IMAGE}!${FLYWAY_IMAGE}!g" ./kubernetes/db/${FILE} | kubectl apply -f -
  done

  for FILE in `ls ./kubernetes/db/ | grep "\-job.yaml$"`;do
    sed -e "s!\${NAMESPACE}!${NAMESPACE}!g" -e "s!\${IMAGE}!${IMAGE}!g" -e "s!\${METADATA_IMAGE}!${METADATA_IMAGE}!g" -e "s!\${FLYWAY_IMAGE}!${FLYWAY_IMAGE}!g" ./kubernetes/db/${FILE} | kubectl delete -f - --ignore-not-found=true
    sed -e "s!\${NAMESPACE}!${NAMESPACE}!g" -e "s!\${IMAGE}!${IMAGE}!g" -e "s!\${METADATA_IMAGE}!${METADATA_IMAGE}!g" -e "s!\${FLYWAY_IMAGE}!${FLYWAY_IMAGE}!g" ./kubernetes/db/${FILE} | kubectl apply -f -
  done

  for FILE in `ls ./kubernetes/etcd/`;do
    sed -e "s!\${NAMESPACE}!${NAMESPACE}!g" -e "s!\${IMAGE}!${IMAGE}!g" -e "s!\${METADATA_IMAGE}!${METADATA_IMAGE}!g" -e "s!\${FLYWAY_IMAGE}!${FLYWAY_IMAGE}!g" ./kubernetes/etcd/${FILE} | kubectl apply -f -
  done

  for FILE in `ls ./kubernetes/openpitrix/ | grep "^openpitrix-"`;do
    sed -e "s!\${NAMESPACE}!${NAMESPACE}!g" -e "s!\${IMAGE}!${IMAGE}!g" -e "s!\${METADATA_IMAGE}!${METADATA_IMAGE}!g" -e "s!\${FLYWAY_IMAGE}!${FLYWAY_IMAGE}!g" ./kubernetes/openpitrix/${FILE} | kubectl apply -f -
  done
fi
if [ "${METADATA}" == "1" ];then
  for FILE in `ls ./kubernetes/openpitrix/metadata/`;do
    sed -e "s!\${NAMESPACE}!${NAMESPACE}!g" -e "s!\${IMAGE}!${IMAGE}!g" -e "s!\${METADATA_IMAGE}!${METADATA_IMAGE}!g" -e "s!\${FLYWAY_IMAGE}!${FLYWAY_IMAGE}!g" ./kubernetes/openpitrix/metadata/${FILE} | kubectl apply -f -
  done
fi

echo "Deployed successfully"


