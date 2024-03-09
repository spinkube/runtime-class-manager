FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY pkg ./pkg
COPY cmd ./cmd
COPY *.go ./
COPY testdata ./testdata

RUN go test ./...

RUN CGO_ENABLED=0 go build -o kwasm-node-installer
RUN /app/kwasm-node-installer -h

FROM scratch
COPY --from=builder /app/kwasm-node-installer /kwasm-node-installer

ENTRYPOINT ["/kwasm-node-installer"]
