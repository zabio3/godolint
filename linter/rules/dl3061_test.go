package rules

import (
	"testing"
)

func TestValidateDL3061(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `RUN apt-get update
FROM ubuntu:latest
`,
			expectedRst: []ValidateResult{
				{line: 1},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
RUN apt-get update
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `ARG VERSION=latest
FROM ubuntu:$VERSION
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `LABEL maintainer="test@example.com"
FROM ubuntu:latest
`,
			expectedRst: []ValidateResult{
				{line: 1},
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3061(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
