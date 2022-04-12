FROM golang:1.16-alpine as builder

RUN apk add --no-cache g++ git
# override git so go can access private repos
# ONLY use this if there is 2nd stage container, otherwise ARG is visible in image history
ARG GH_TOKEN
RUN git config --global url."https://$GH_TOKEN:@github.com/".insteadOf "https://github.com/"

WORKDIR /sgn-v2
ADD go.mod go.sum /sgn-v2/
ENV GOPRIVATE github.com/celer-network/cbridge-flow
RUN go mod download

ADD . /sgn-v2
ENV GOPRIVATE github.com/celer-network/cbridge-flow
RUN go build -o /sgn-v2/bin/executor ./executor/main

FROM alpine:latest
VOLUME /executor/env
WORKDIR /executor/env
COPY --from=builder /sgn-v2/bin/executor /usr/local/bin
CMD ["/bin/sh", "-c", "executor start --test --loglevel debug --home /executor/env 2> /executor/env/app.log"]