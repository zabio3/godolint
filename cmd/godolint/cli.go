package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/zabio3/godolint/linter"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagsError
	ExitCodeNoExistError
	ExitCodeFileError
	ExitCodeAstParseError
	ExitCodeLintCheckError
)

const name = "godolint"

const usage = `Usage: godolint <Dockerfile>
godolint is a Dockerfile linter command line tool that helps you build best practice Docker images.
`

// CLI represents CLI interface
type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) run(args []string) int {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(cli.outStream, usage)
	}

	// ToDo Set necessary flags (--ignore-rule)
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	// The Dockerfile to be analyzed must be the last.
	if len(args) < 2 {
		_, _ = fmt.Fprintf(cli.errStream, "Please provide a Dockerfile\n")
		return ExitCodeNoExistError
	}

	file := args[len(args)-1]
	f, err := os.Open(file)
	if err != nil {
		_, _ = fmt.Fprintf(cli.errStream, "%s\n", err)
		return ExitCodeFileError
	}

	r, err := parser.Parse(f)
	if err != nil {
		_, _ = fmt.Fprintf(cli.errStream, "%s\n", err)
		return ExitCodeAstParseError
	}

	rst, err := linter.Analize(r.AST)
	if err != nil {
		_, _ = fmt.Fprintf(cli.errStream, "%s\n", err)
		return ExitCodeLintCheckError
	}

	_, _ = fmt.Fprintf(cli.outStream, "%s\n", rst)
	return ExitCodeOK
}
