package rules

import (
	"testing"
)

func TestValidateDL3002(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM golang:latest

USER root
WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
`,
			file:        "DL3002_Dockerfile",
			expectedRst: []string{"DL3002_Dockerfile:3 DL3002 Last USER should not be root\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM golang:latest

USER root
WORKDIR /go
ADD . /go
USER zabio3

CMD ["go", "run", "main.go"]
`,
			file:        "DL3002_Dockerfile_2",
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3002(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
