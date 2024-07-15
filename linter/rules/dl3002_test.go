package rules

import (
	"testing"
)

func TestValidateDL3002(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM golang:latest

USER root
WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
`,
			expectedRst: []ValidateResult{
				{line: 3},
			},
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
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3002(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
