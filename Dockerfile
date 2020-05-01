FROM golang:1.14.2-alpine3.11

RUN apk update && \
    apk upgrade && \
    apk add git && \
    apk add openssh-client

EXPOSE 8000