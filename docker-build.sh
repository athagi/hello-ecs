#!/bin/bash
COMMIT_HASH=${1:-none}
IS_DEVEL=${2:-true} # devel or latest
TAG="latest"
APP_NAME="my-app"

cd "$(dirname "$0")" || exit

if [ "${IS_DEVEL}" = "true" ]; then
  TAG="devel"
fi

DOCKER_BUILDKIT=1 docker build . --file Dockerfile --tag "${APP_NAME}":"${COMMIT_HASH}" --tag "${APP_NAME}":"${TAG}" --build-arg COMMIT_HASH="${COMMIT_HASH}"

