# Pull base image.
FROM ubuntu:14.04

# tell the port number the container should expose
EXPOSE 9090

RUN mkdir -p /go/src/app/uuid-server

ADD /bin /go/src/app/uuid-server

WORKDIR /go/src/app/uuid-server

ENTRYPOINT ./server
