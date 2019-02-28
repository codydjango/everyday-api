#!/bin/bash

BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )";
ROOT_DIR="$(dirname $BIN_DIR)";
SRC_DIR="$ROOT_DIR/src";
ENVIRONMENT="production";
VIRTUAL_HOST="everyday.invisiblehands.ca";
LETSENCRYPT_HOST="everyday.invisiblehands.ca";
LETSENCRYPT_EMAIL="cody@invisiblehands.ca";

docker run --detach \
    --name goserver \
    --env "ENVIRONMENT=${ENVIRONMENT}" \
    --env "VIRTUAL_HOST=${VIRTUAL_HOST}" \
    --env "LETSENCRYPT_HOST=${LETSENCRYPT_HOST}" \
    --env "LETSENCRYPT_EMAIL=${LETSENCRYPT_EMAIL}" \
    goserver;