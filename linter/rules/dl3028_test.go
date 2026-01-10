package rules

import (
	"testing"
)

func TestValidateDL3028(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM ruby:2.5
RUN gem install bundler
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ruby:2.5
RUN gem install bundler:2.1.4
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ruby:2.5
RUN gem install bundler -v 2.1.4
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM ruby:2.5
RUN gem install bundler:2.1.4 && gem install unpinned
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3028(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
