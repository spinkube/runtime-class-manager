FROM alpine:3.21.2
WORKDIR /
COPY ./bin/manager /manager

ENTRYPOINT ["/manager"]
