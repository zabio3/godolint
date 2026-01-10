package rules

import (
	"testing"
)

func TestValidateDL3043(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM ubuntu:latest
ONBUILD FROM alpine
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
ONBUILD MAINTAINER test@example.com
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
ONBUILD ONBUILD RUN echo test
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
ONBUILD RUN apt-get update
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
ONBUILD ADD . /app
ONBUILD RUN make /app
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

		gotRst, gotErr := validateDL3043(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
