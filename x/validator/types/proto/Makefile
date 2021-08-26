proto:
	buf protoc --gocosmos_out=plugins=interfacetype+grpc:. types.proto
	mv github.com/celer-network/sgn-v2/x/validator/types/types.pb.go ../
	rm -r github.com