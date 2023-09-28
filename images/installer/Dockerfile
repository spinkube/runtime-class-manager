FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY pkg ./pkg
COPY cmd ./cmd
COPY *.go ./

RUN go build -o kwasm-node-installer

CMD ["/app/kwasm-node-installer", "install"]
