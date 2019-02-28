#!/bin/bash

BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )";
ROOT_DIR="$(dirname $BIN_DIR)";
SRC_DIR="$ROOT_DIR/src";
VIRTUAL_HOST=everyday.invisiblehands.ca

docker run --rm -p 3001:3001 -it --name goserver -e VIRTUAL_HOST=${VIRTUAL_HOST} --mount source=${SRC_DIR},destination=/go/app/src,type=bind goserver