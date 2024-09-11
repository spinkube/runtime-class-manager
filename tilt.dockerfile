FROM alpine@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d
WORKDIR /
COPY ./bin/manager /manager

ENTRYPOINT ["/manager"]
