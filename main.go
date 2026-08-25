package main

import (
	"os"

	"weber-br/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
