package rules

import (
	"testing"
)

func TestValidateDL3054(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.licenses="MIT"
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.licenses="Apache-2.0"
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.licenses="CustomLicense"
`,
			expectedRst: []ValidateResult{
				{line: 2, addMsg: "Label value is not a valid SPDX license: org.opencontainers.image.licenses"},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
LABEL org.opencontainers.image.licenses="GPL-3.0"
`,
			expectedRst: nil,
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

		gotRst, gotErr := validateDL3054(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
