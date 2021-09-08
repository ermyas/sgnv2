package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	sgndimpl "github.com/celer-network/sgn-v2/cmd/sgnd/impl"
	"github.com/spf13/cobra/doc"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory] \n", os.Args[0])
		os.Exit(1)
	}
	path := os.Args[1]

	sgndPath := filepath.Join(path, "sgnd")
	os.RemoveAll(sgndPath)
	os.Mkdir(sgndPath, 0755)
	rootCmd, _ := sgndimpl.NewRootCmd()
	err := doc.GenMarkdownTree(rootCmd, sgndPath)
	if err != nil {
		log.Fatal(err)
	}
}
