package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	cases := []struct {
		command           string
		expectedOutStream string
		expectedErrStream string
		expectedExitCode  int
	}{
		{
			command:           "godolint",
			expectedOutStream: "",
			expectedErrStream: "Please provide a Dockerfile\n",
			expectedExitCode:  ExitCodeNoExistError,
		},
		{
			command: "godolint -h",
			expectedOutStream: `godolint - Dockerfile Linter written in Golang

Usage: godolint [--ignore RULECODE]
  Lint Dockerfile for errors and best practices

Available options:
  --ignore RULECODE        A rule to ignore. If present, the ignore list in the
                           config file is ignored
`,
			expectedErrStream: "flag: help requested\n",
			expectedExitCode:  ExitCodeParseFlagsError,
		},
		{
			command:           "godolint testdata/no-file",
			expectedOutStream: "",
			expectedErrStream: "open testdata/no-file: no such file or directory\n",
			expectedExitCode:  ExitCodeFileError,
		},
		{
			command:           "godolint ../testdata/OK_Dockerfile",
			expectedOutStream: "",
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
		{
			command:           "godolint ../testdata/DL3000_Dockerfile",
			expectedOutStream: "../testdata/DL3000_Dockerfile:3 DL3000 Use absolute WORKDIR\n",
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
		{
			command:           "godolint ../testdata/DL3001_Dockerfile",
			expectedOutStream: "../testdata/DL3001_Dockerfile:6 DL3001 For some bash commands it makes no sense running them in a Docker container like `ssh`, `vim`, `shutdown`, `service`, `ps`, `free`, `top`, `kill`, `mount`, `ifconfig`\n",
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
		{
			command:           "godolint --ignore DL3001 ../testdata/DL3001_Dockerfile",
			expectedOutStream: "",
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
		{
			command:           "godolint ../testdata/DL3002_Dockerfile",
			expectedOutStream: "../testdata/DL3002_Dockerfile:3 DL3002 Last USER should not be root\n",
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
		{
			command:           "godolint ../testdata/DL3003_Dockerfile",
			expectedOutStream: "../testdata/DL3003_Dockerfile:6 DL3003 Use WORKDIR to switch to a directory\n",
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
		{
			command:           "godolint --ignore DL3009 ../testdata/DL3004_Dockerfile",
			expectedOutStream: "../testdata/DL3004_Dockerfile:3 DL3004 Do not use sudo as it leads to unpredictable behavior. Use a tool like gosu to enforce root.\n",
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
		{
			command:           "godolint ../testdata/MaxScanSize_File",
			expectedOutStream: "",
			expectedErrStream: "dockerfile line greater than max allowed size of 65535\n",
			expectedExitCode:  ExitCodeAstParseError,
		},
		{
			command:           "godolint --ignore NO_RULE ../testdata/OK_Dockerfile",
			expectedOutStream: "",
			expectedErrStream: "no exist rule specified by ignore flag: NO_RULE\n",
			expectedExitCode:  ExitCodeLintCheckError,
		},
	}

	for i, tc := range cases {
		outStream := new(bytes.Buffer)
		errStream := new(bytes.Buffer)

		cli := CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(tc.command, " ")

		if got := cli.run(args); got != tc.expectedExitCode {
			t.Errorf("#%d %q exits: want: %d, got: %d", i, tc.command, tc.expectedExitCode, got)
		}

		if got := outStream.String(); got != tc.expectedOutStream {
			t.Errorf("#%d Unexpected outStream has returned: want: %s, got: %s", i, tc.expectedOutStream, got)
		}

		if got := errStream.String(); got != tc.expectedErrStream {
			t.Errorf("#%d Unexpected errStream has returned: want:%s, got:%s", i, tc.expectedErrStream, got)
		}

		cleanup(t)
	}
}

func cleanup(t *testing.T) {
	t.Helper()
}
