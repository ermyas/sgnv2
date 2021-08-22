package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/celer-network/sgn-v2/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory] \n", os.Args[0])
		os.Exit(1)
	}
	path := os.Args[1]

	sgncliPath := filepath.Join(path, "sgncli")
	os.RemoveAll(sgncliPath)
	os.Mkdir(sgncliPath, 0755)
	err := doc.GenMarkdownTree(cmd.GetSgncliExecutor().Command, sgncliPath)
	if err != nil {
		log.Fatal(err)
	}
}
