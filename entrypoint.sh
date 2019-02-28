#!/bin/bash
if [ -z "$ENVIRONMENT" ] 
then
    ENVIRONMENT="development"
fi

echo "entrypoint environment: ${ENVIRONMENT}"

if [ "$ENVIRONMENT" = "production" ]
then
    echo "starting production server"
    go build -o ../build/server *.go && ../build/server
else
    echo "starting development server"
    reflex -s -r \.go$ -- go run *.go
fi