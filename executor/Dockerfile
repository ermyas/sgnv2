FROM golang:1.16-alpine as builder

RUN apk add --no-cache g++ git
# override git so go can access private repos
# ONLY use this if there is 2nd stage container, otherwise ARG is visible in image history
ARG GH_TOKEN
RUN git config --global url."https://$GH_TOKEN:@github.com/".insteadOf "https://github.com/"

RUN git clone https://github.com/celer-network/sgn-v2-executor.git && git -C sgn-v2-executor checkout 0c64a93d7f6b8af0edf7e24565a7ad076b66cf1c

WORKDIR sgn-v2-executor
RUN go mod download
RUN go build -o /sgn-v2-executor/bin/executor ./cmd/main

FROM alpine:latest
VOLUME /executor/env
WORKDIR /executor/env
COPY --from=builder /sgn-v2-executor/bin/executor /usr/local/bin
CMD ["/bin/sh", "-c", "executor start --test --loglevel debug --home /executor/env > /executor/env/app.log 2>&1"]