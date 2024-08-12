package rules

import (
	"testing"
)

func TestValidateDL3013(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM python:3.4
RUN pip install django
RUN pip install https://github.com/Banno/carbon/tarball/0.9.x-fix-events-callback
`,
			expectedRst: []ValidateResult{
				{line: 2},
				{line: 3},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM python:3.4
RUN pip install django && pip install https://github.com/Banno/carbon/tarball/0.9.x-fix-events-callback
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM python:3.4
ARG YAML_LINT_VERSION=v1.26.3

RUN pip install --no-cache-dir yamllint=="${YAML_LINT_VERSION:-v1.26.3}"
RUN pip install django && pip install https://github.com/Banno/carbon/tarball/0.9.x-fix-events-callback
`,
			expectedRst: []ValidateResult{
				{line: 5},
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3013(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
