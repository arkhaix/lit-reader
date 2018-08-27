#! /bin/bash

DO_COMPOSE=0
DO_GKE=0

# Check flags
while getopts "cg" o; do
  case "${o}" in 
    c)
      DO_COMPOSE=1
      ;;
    g)
      DO_GKE=1
      ;;
  esac
done
shift $(($OPTIND-1))

# Set up
DOCKER_IMAGE=${1}
DOCKER_COMPOSE_TAG=docker-compose_${DOCKER_IMAGE}
GKE_TAG=gcr.io/arkhaix-lit-reader/${DOCKER_IMAGE}

# Build
docker build -t ${DOCKER_IMAGE} -f ./build/${DOCKER_IMAGE}.Dockerfile .

# Tag docker-compose
if [[ ! ${DO_COMPOSE} -eq 0 ]]; then
  echo "tagging for docker-compose: ${DOCKER_COMPOSE_TAG}"
  docker tag ${DOCKER_IMAGE} ${DOCKER_COMPOSE_TAG}
fi

# Tag and push gke
if [[ ! ${DO_GKE} -eq 0 ]]; then
  echo "tagging for gke: ${GKE_TAG}"
  docker tag ${DOCKER_IMAGE} ${GKE_TAG}

  echo "pushing to gke ${GKE_TAG}"
  docker push ${GKE_TAG}
fi
