PROOT=../..
protoc -I$PROOT -I/usr/local/include --go_out=plugins=grpc:$PROOT --grpc-gateway_out=$PROOT $PROOT/proto/sgn/cbridge/v1/gateway.proto
git add $PROOT/gateway/webapi/