FROM golang:1.16-alpine as builder

RUN apk add --no-cache g++

WORKDIR /sgn-v2
ADD go.mod go.sum /sgn-v2/
RUN go mod download

ADD . /sgn-v2
RUN go build -o /sgn-v2/bin/gateway ./gateway/main/main.go

FROM alpine:latest
VOLUME /sgn-v2/env
WORKDIR /sgn-v2/env
EXPOSE 18081 18082 20000
COPY --chmod=0777 --from=builder /sgn-v2/bin/gateway /usr/local/bin
CMD gateway -home=/sgn-v2/env/sgnd 2> /sgn-v2/env/sgnd/app.log
STOPSIGNAL SIGTERM