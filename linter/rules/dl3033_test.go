package rules

import (
	"testing"
)

func TestValidateDL3033(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM centos
RUN yum install python
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM centos
RUN yum install -y python-2.7.5
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM centos
RUN yum install -y --setopt=tsflags=nodocs python-2.7.5 ruby-2.0.0
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

		gotRst, gotErr := validateDL3033(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
