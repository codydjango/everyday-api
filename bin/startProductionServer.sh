#!/bin/bash

BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )";
ROOT_DIR="$(dirname $BIN_DIR)";
SRC_DIR="$ROOT_DIR/src";
VIRTUAL_HOST="everyday.invisiblehands.ca"
ENVIRONMENT="production"

docker run --rm -it --name goserver -e VIRTUAL_HOST=${VIRTUAL_HOST} -e ENVIRONMENT=${ENVIRONMENT} goserver