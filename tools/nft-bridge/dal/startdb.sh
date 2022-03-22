#!/bin/sh
# $1 is db folder

set -e
# use 18080 to avoid port conflict w/ grpc svr
cockroach start-single-node --insecure --listen-addr=localhost:26257 --http-addr=localhost:18080 --store=path=$1 &
# Ensure CockroachDB is up
sleep 2
nc -v -w 5 -z localhost 26257
cat schema.sql | cockroach sql --insecure
