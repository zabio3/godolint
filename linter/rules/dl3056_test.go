package rules

import (
	"testing"
)

func TestValidateDL3056(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.version="1.0.0"
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.version="1.0.0-alpha.1"
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.version="1.0.0+20230115"
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.version="v1.0"
`,
			expectedRst: []ValidateResult{
				{line: 2, addMsg: "Label value is not a valid semantic version: org.opencontainers.image.version"},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.version="1"
`,
			expectedRst: []ValidateResult{
				{line: 2, addMsg: "Label value is not a valid semantic version: org.opencontainers.image.version"},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL description="My App"
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

		gotRst, gotErr := validateDL3056(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
