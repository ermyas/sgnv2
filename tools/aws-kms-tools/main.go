package main

import (
	"os"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/tools/aws-kms-tools/impl"
)

func main() {
	rootCmd := impl.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
