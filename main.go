package main

import (
	"github.com/zabio3/godolint/cmd"
	"os"
)

func main() {
	cli := &cmd.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
