FROM golang:1.16-alpine as builder

RUN apk add --no-cache g++

WORKDIR /sgn-v2
ADD go.mod go.sum /sgn-v2/
RUN go mod download

ADD . /sgn-v2
RUN go build -o /sgn-v2/bin/executor ./executor/main

FROM alpine:latest
VOLUME /executor/env
WORKDIR /executor/env
COPY --from=builder /sgn-v2/bin/executor /usr/local/bin
CMD ["/bin/sh", "-c", "executor -test -loglevel=debug -home=/executor/env 2> /executor/env/app.log"]