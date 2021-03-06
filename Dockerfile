FROM golang:alpine
RUN apk add git bash build-base
RUN go get github.com/cespare/reflex
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get github.com/ethereum/go-ethereum/crypto
RUN go get github.com/ethereum/go-ethereum/common
RUN go get github.com/ethereum/go-ethereum/common/hexutil
RUN go get github.com/dgrijalva/jwt-go

SHELL ["/bin/bash", "-c"]
ENV PORT=3001
EXPOSE 3001

ADD ./src ./app/src
ADD ./data ./app/data
ADD ./entrypoint.sh /go/app/entrypoint.sh

WORKDIR /go/app/src
ENTRYPOINT ["/go/app/entrypoint.sh"]