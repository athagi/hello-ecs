#!/bin/bash
COMMIT_HASH=${1:-none}
IS_DEVEL=${2:-true} # devel or latest
TAG="latest"

if [ "${IS_DEVEL}" = "true" ]; then
  TAG="devel"
fi

DOCKER_BUILDKIT=1 docker build --tag "${COMMIT_HASH}" --tag "${TAG}" .

