package rules

import (
	"testing"
)

func TestValidateDL4006(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `RUN wget -O - https://some.site | wc -l > /number
`,
			expectedRst: []ValidateResult{
				{line: 1},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN wget -O - https://some.site | wc -l > /number
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN yq eval '.".docker".script[1] | explode(.)' /base.yml > /usr/local/bin/entrypoint
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

		gotRst, gotErr := validateDL4006(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
