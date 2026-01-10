package rules

import (
	"testing"
)

func TestValidateDL3055(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.revision="abcdef1234567890abcdef1234567890abcdef12"
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.revision="abcdef12345"
`,
			expectedRst: []ValidateResult{
				{line: 2, addMsg: "Label value is not a valid git hash: org.opencontainers.image.revision"},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.revision="xyz1234567890abcdef1234567890abcdef1234"
`,
			expectedRst: []ValidateResult{
				{line: 2, addMsg: "Label value is not a valid git hash: org.opencontainers.image.revision"},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL version="1.0.0"
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

		gotRst, gotErr := validateDL3055(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
