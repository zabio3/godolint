package rules

import (
	"testing"
)

func TestValidateDL3003(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM golang:latest

WORKDIR /go
ADD . /go

RUN cd /usr/src/app && git clone git@github.com:zabio3/godolint.git /usr/src/app

CMD ["go", "run", "main.go"]
`,
			file:        "DL3003_Dockerfile",
			expectedRst: []string{"DL3003_Dockerfile:6 DL3003 Use WORKDIR to switch to a directory\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3003(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
