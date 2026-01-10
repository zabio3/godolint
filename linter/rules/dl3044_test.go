package rules

import (
	"testing"
)

func TestValidateDL3044(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM ubuntu:latest
ENV PATH=$PATH:/usr/local/bin
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
ENV MY_VAR=${MY_VAR}:additional
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
ENV VAR1=value1
ENV VAR2=$VAR1
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
ENV MY_VAR=value
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

		gotRst, gotErr := validateDL3044(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
