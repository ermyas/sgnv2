package main

import (
	"os"

	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/gateway"
)

func main() {
	encodingConfig := app.MakeEncodingConfig()
	gateway.InitGateway(os.ExpandEnv("$HOME/.sgnd"), encodingConfig.Amino, encodingConfig.Codec, encodingConfig.InterfaceRegistry, true, "127.0.0.1:26257")
}
