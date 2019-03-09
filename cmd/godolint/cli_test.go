package main

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
			command: "godolint -h",
			expectedOutStream: `Usage: godolint <Dockerfile>
godolint is a Dockerfile linter command line tool that helps you build best practice Docker images.
`,
			expectedErrStream: "",
			expectedExitCode:  ExitCodeParseFlagsError,
		},
	}

	for i, tc := range cases {
		outStream := new(bytes.Buffer)
		errStream := new(bytes.Buffer)

		cli := CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(tc.command, " ")

		if got := cli.Run(args); got != tc.expectedExitCode {
			t.Errorf("#%d %q exits with %d, want %d", i, tc.command, got, tc.expectedExitCode)
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
