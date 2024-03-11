FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o kwasm-node-installer cmd/node-installer/main.go
RUN /app/kwasm-node-installer -h

FROM scratch
COPY --from=builder /app/kwasm-node-installer /kwasm-node-installer

ENTRYPOINT ["/kwasm-node-installer"]
