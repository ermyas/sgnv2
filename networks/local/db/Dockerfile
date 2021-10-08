FROM cockroachdb/cockroach:latest

EXPOSE 26257
ENTRYPOINT [ "cockroach", "start-single-node", "--insecure", "--listen-addr=:26257" ]

STOPSIGNAL SIGTERM