FROM golang:alpine
RUN apk add git bash
RUN go get github.com/cespare/reflex
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors

SHELL ["/bin/bash", "-c"]
ENV PORT=3001

ADD ./src ./app/src
ADD ./entrypoint.sh /go/app/entrypoint.sh

WORKDIR /go/app/src
ENTRYPOINT ["/go/app/entrypoint.sh"]