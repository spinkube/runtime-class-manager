FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY pkg ./pkg
COPY cmd ./cmd
COPY *.go ./

RUN go build -o kwasm-node-installer

FROM scratch
COPY --from=builder /app/kwasm-node-installer /kwasm-node-installer

ENTRYPOINT ["/kwasm-node-installer"]
