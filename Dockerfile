FROM golang:1.16-alpine as builder

RUN apk add --no-cache g++ musl-dev linux-headers leveldb-dev

WORKDIR /sgn-v2
ADD go.mod go.sum /sgn-v2/
RUN go mod download

ADD . /sgn-v2
RUN go build -tags "cleveldb" -o /sgn-v2/bin/sgnd ./cmd/sgnd

FROM alpine:latest
RUN apk add leveldb
VOLUME /sgn-v2/env
WORKDIR /sgn-v2/env
EXPOSE 26656 26657 1317 9090 9091
COPY --from=builder /sgn-v2/bin/sgnd /usr/local/bin
CMD ["/bin/sh", "-c", "sgnd start --home /sgn-v2/env/sgnd 2> /sgn-v2/env/sgnd/app.log > /sgn-v2/env/sgnd/tendermint.log"]
STOPSIGNAL SIGTERM