FROM alpine:3.21.0
WORKDIR /
COPY ./bin/manager /manager

ENTRYPOINT ["/manager"]
