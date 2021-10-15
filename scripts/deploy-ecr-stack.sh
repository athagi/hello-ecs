#!/bin/bash
STACK_NAME=${1:-ecr-test-stack}
REPO_NAME=${2:-test}

cd `dirname $0`

rain deploy ../ecr.yaml --params ECRRepoName="${REPO_NAME}" "${STACK_NAME}" -y
