package main

import (
	"github.com/zabio3/godolint/cmd"
	"os"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.run(os.Args))
}
