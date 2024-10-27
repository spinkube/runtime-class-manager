FROM alpine:3.20.3
WORKDIR /
COPY ./bin/manager /manager

ENTRYPOINT ["/manager"]
