#!/bin/bash
STACK_NAME=${1:-ecr-test-stack}
REPO_NAME=${2:-test}

aws cloudformation deploy --template-file ecr.yaml --stack-name "${STACK_NAME}" --parameter-override ECRRepoName="${REPO_NAME}"