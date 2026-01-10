package rules

import (
	"testing"
)

func TestValidateDL3057(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM ubuntu:latest
RUN apt-get update
CMD ["nginx"]
`,
			expectedRst: []ValidateResult{
				{line: 1, addMsg: "No HEALTHCHECK instruction found"},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
RUN apt-get update
HEALTHCHECK CMD curl -f http://localhost/ || exit 1
CMD ["nginx"]
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ubuntu:latest
HEALTHCHECK --interval=30s --timeout=3s CMD curl -f http://localhost/ || exit 1
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

		gotRst, gotErr := validateDL3057(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
