// Package cmd provides command line tool to analyzer dockerfile.
package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

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

const version = "1.0.3"

const usage = `godolint - Dockerfile linter written in Golang

Usage: godolint [--ignore RULECODE]
  Lint Dockerfile for errors and best practices

Available options:
  --ignore RULECODE	A rule to ignore. If present, the ignore list in the
			config file is ignored
  --trusted-registry REGISTRY (e.g. docker.io)
			A docker registry to allow to appear in FROM instructions

  --help	-h	Print this help message and exit.
  --version	-v	Print the version information
`

// CLI represents CLI interface.
type CLI struct {
	OutStream, ErrStream io.Writer
}

type sliceString []string

func (ss *sliceString) String() string {
	return fmt.Sprintf("%s", *ss)
}

func (ss *sliceString) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}

// Run it takes Dockerfile as an argument and applies it to analyzer to standard output.
func (cli *CLI) Run(args []string) int {
	var ignoreRules sliceString
	var isVersion bool
	var trustedRegistries sliceString

	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(cli.OutStream, usage)
	}

	flags.Var(&ignoreRules, "ignore", "Set ignore strings")
	flags.BoolVar(&isVersion, "version", false, "version")
	flags.BoolVar(&isVersion, "v", false, "version")
	flags.Var(&trustedRegistries, "trusted-registry", "Docker registries to allow to appear in FROM instructions")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprint(cli.ErrStream, err)
		return ExitCodeParseFlagsError
	}

	if isVersion {
		fmt.Fprintf(cli.OutStream, "godolint version %v\n", version)
		return ExitCodeOK
	}

	length := len(args)
	// The Dockerfile to be analyzed must be the last.
	if length < 2 {
		fmt.Fprintf(cli.ErrStream, "Please provide a Dockerfile\n")
		return ExitCodeNoExistError
	}

	file := args[length-1]
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprint(cli.ErrStream, err)
		return ExitCodeFileError
	}

	r, err := parser.Parse(f)
	if err != nil {
		fmt.Fprint(cli.ErrStream, err)
		return ExitCodeAstParseError
	}

	analyzer := linter.NewAnalyzer(ignoreRules, trustedRegistries)
	rst, err := analyzer.Run(r.AST)
	if err != nil {
		fmt.Fprint(cli.ErrStream, err)
		return ExitCodeLintCheckError
	}

	rst = sort.StringSlice(rst)
	var output string
	for i := range rst {
		// ends of each strings have "\n"
		output = output + rst[i]
	}
	fmt.Fprint(cli.OutStream, output)
	return ExitCodeOK
}
