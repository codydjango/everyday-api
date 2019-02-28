#!/bin/bash
if [ -z "$ENVIRONMENT" ] 
then
    ENVIRONMENT="development"
fi

echo "entrypoint environment: ${ENVIRONMENT}"

if [ "$ENVIRONMENT" = "production" ]
then
    reflex -s -r \.go$ -- go run *.go
else
    go run *.go
fi