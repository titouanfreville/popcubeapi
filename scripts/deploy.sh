#!/bin/bash
#
# Popcube Release branch and deployement
# MAINTAINER - Clément LE CORRE
DEPLOY=false
REPO="registry.le-corre.eu:5000"

function usage(){
    echo "Popcube Release branch and deployement"
    echo ""
    echo ""
    echo -e "$0"
    echo -e "\t-h --help"
    echo -e "\t--deploy=$DEPLOY"
    echo -e "\t--branch=$BRANCH"
    echo -e "\t--tag=$TAG"
    echo ""
}
function build_images() {
  # First args is tags
  docker build --no-cache -t ${REPO}/popcubedocs:$1 -f docker/slateserver.Dockerfile .
  docker build --no-cache -t ${REPO}/popcubedb:$1 -f docker/database.Dockerfile .
  docker build --no-cache -t ${REPO}/popcubeapi:$1 -f docker/gobase.Dockerfile .
}
function push_images() {
  # First args is tags
  docker push ${REPO}/popcubedocs:$1
  docker push ${REPO}/popcubedb:$1
  docker push ${REPO}/popcubeapi:$1
}
function deployement() {

curl -X POST --header 'Content-Type: application/json' \
 --header "X-AUTH-TOKEN: ${DEPLOY_TOKEN}" -d \
 "{
   \"Image\": \"${REPO}/popcubedocs:$1\",
   \"Env\": [
     \"VIRTUAL_NETWORK=nginx-proxy\",
     \"VIRTUAL_HOST=docs-alpha.popcube.xyz\",
     \"VIRTUAL_PORT=4567\"
   ],
   \"Hostname\": \"popcube_alpha_docs\" }" \
    http://${DEPLOY_URL}/deploy;

curl -X POST --header 'Content-Type: application/json' \
  --header "X-AUTH-TOKEN: ${DEPLOY_TOKEN}" -d \
  "{
    \"Image\": \"${REPO}/popcubedb:$1\",
    \"Env\": [
      \"MYSQL_PASSWORD=test\",
      \"MYSQL_ROOT_PASSWORD=popcube_dev\",
      \"MYSQL_USER=test_user\",
      \"MYSQL_DATABASE=popcube_test\"
    ],
    \"Hostname\": \"popcube_alpha_database\" }" \
     http://${DEPLOY_URL}/deploy;

# API BACK
curl -X POST --header 'Content-Type: application/json' \
 --header "X-AUTH-TOKEN: ${DEPLOY_TOKEN}" -d \
 "{
   \"Image\": \"${REPO}/popcubeapi:$1\",
   \"Env\": [
     \"VIRTUAL_NETWORK=nginx-proxy\",
     \"VIRTUAL_HOST=api-alpha.popcube.xyz\",
     \"VIRTUAL_PORT=3000\"
   ],
    \"HostConfig\": {
      \"Links\": [
                  \"/popcube_alpha_database:/popcube_alpha_api/database\"
              ]
    },
   \"Hostname\": \"popcube_alpha_api\" }" \
    http://${DEPLOY_URL}/deploy;
}
if [ "$#" -eq 0 ]; then
    usage
    exit 0
fi
while [ "$1" != "" ]; do
    PARAM=`echo $1 | awk -F= '{print $1}'`
    VALUE=`echo $1 | awk -F= '{print $2}'`
    case $PARAM in
        -h | --help)
            usage
            exit
            ;;
        --tag)
            TAG=$VALUE
            ;;
        --branch)
            BRANCH=$VALUE
            ;;
        --deploy)
            DEPLOY=$VALUE
            ;;
        *)
            echo "ERROR: unknown parameter \"$PARAM\""
            usage
            exit 1
            ;;
    esac
    shift
done


if [ ${TAG+x} ] && [ ${BRANCH+x} ];
then
  echo "TAG and BRANCH is not compatible :("
  exit 1
fi
if [ ${TAG+x} ];
then
  echo "TAG is set to '$TAG'";
  build_images "$TAG"
  push_images "$TAG"
elif [ ${BRANCH+x} ];
then
  echo "BRANCH is set to '$BRANCH'";
  if [ ${BRANCH} = "master" ];
  then
    echo "branch develop : OK"
    build_images "master"
    push_images "master"
    build_images "latest"
    push_images "latest"
  elif [ ${BRANCH} = "develop" ];
  then
    echo "branch develop : OK"
    build_images "dev"
    push_images "dev"
  else
    echo "Release branch not found.."
  fi
fi
if [ ${DEPLOY} = "true" ];
then
  if [ ${TAG+x} ];
  then
    deployement "${TAG}"
  else
    echo "tag branch not found.."
  fi
else
  echo "Deployement is disable"
fi
