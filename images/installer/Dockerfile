FROM golang:1.23@sha256:ad5c126b5cf501a8caef751a243bb717ec204ab1aa56dc41dc11be089fafcb4f AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o kwasm-node-installer ./cmd/node-installer
RUN /app/kwasm-node-installer -h

FROM scratch
COPY --from=builder /app/kwasm-node-installer /kwasm-node-installer

ENTRYPOINT ["/kwasm-node-installer"]
