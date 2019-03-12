package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/zabio3/godolint/linter"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK = iota
	ExitCodeParseFlagsError
	ExitCodeNoExistError
	ExitCodeFileError
	ExitCodeAstParseError
	ExitCodeLintCheckError
)

const name = "godolint"

const usage = `godolint - Dockerfile Linter written in Golang

Usage: godolint [--ignore RULECODE]
  Lint Dockerfile for errors and best practices

Available options:
  --ignore RULECODE        A rule to ignore. If present, the ignore list in the
                           config file is ignored
`

// CLI represents CLI interface
type CLI struct {
	outStream, errStream io.Writer
}

type sliceString []string

func (ss *sliceString) String() string {
	return fmt.Sprintf("%s", *ss)
}

func (ss *sliceString) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}

func (cli *CLI) run(args []string) int {
	var ignoreString sliceString

	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(cli.outStream, usage)
	}

	flags.Var(&ignoreString, "ignore", "Set ignore strings")

	if err := flags.Parse(args[1:]); err != nil {
		_, _ = fmt.Fprintf(cli.errStream, "%s\n", err)
		return ExitCodeParseFlagsError
	}

	length := len(args)
	// The Dockerfile to be analyzed must be the last.
	if length < 2 {
		_, _ = fmt.Fprintf(cli.errStream, "Please provide a Dockerfile\n")
		return ExitCodeNoExistError
	}

	file := args[length-1]
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

	rst, err := linter.Analize(r.AST, file, ignoreString)
	if err != nil {
		_, _ = fmt.Fprintf(cli.errStream, "%s\n", err)
		return ExitCodeLintCheckError
	}

	_, _ = fmt.Fprintf(cli.outStream, strings.Trim(fmt.Sprintf("%s", rst), "[]"))
	return ExitCodeOK
}
