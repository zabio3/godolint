package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagsError
)

const Name = "godolint"

const Usage = `Usage: godolint <Dockerfile>
godolint is a Dockerfile linter command line tool that helps you build best practice Docker images.
`

// CLI represents CLI interface
type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprintf(cli.outStream, Usage)
	}

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	return ExitCodeOK
}
