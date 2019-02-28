#!/bin/bash

BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )";
ROOT_DIR="$(dirname ${BIN_DIR})";
DOCKERFILE="${ROOT_DIR}/Dockerfile";

docker build \
    --tag goserver \
    --file $DOCKERFILE \
    $ROOT_DIR;