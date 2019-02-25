FROM golang:alpine
RUN apk add git bash
RUN go get github.com/cespare/reflex
SHELL ["/bin/bash", "-c"]
ENV PORT=3001w

ADD ./src ./app/src
ADD ./entrypoint.sh /go/app/entrypoint.sh

WORKDIR /go/app/src
ENTRYPOINT ["/go/app/entrypoint.sh"]