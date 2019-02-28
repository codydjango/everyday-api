#!/bin/bash

BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )";
ROOT_DIR="$(dirname $BIN_DIR)";
SRC_DIR="$ROOT_DIR/src";
ENVIRONMENT="development";

docker run \
    --rm \
    --name goserver \
    --publish 3001:3001 \
    --env "ENVIRONMENT=${ENVIRONMENT}" \
    --mount source=${SRC_DIR},destination=/go/app/src,type=bind \
    goserver;