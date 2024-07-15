package rules

import (
	"testing"
)

func TestValidateDL3015(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM busybox
RUN apt-get install -y python=2.7
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM busybox
RUN apt-get install -y python=2.7 && RUN apt-get install -y ruby
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{ // Has two && but only returns one error for the whole line
			dockerfileStr: `FROM busybox
RUN apt-get install -y python=2.7 && RUN apt-get install -y ruby && RUN apt-get install -y nodejs
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{ // Three install commands, first one is missing --no-install-recommends
			dockerfileStr: `FROM busybox
RUN apt-get install -y python=2.7 && RUN apt-get install --no-install-recommends -y ruby && RUN apt-get install --no-install-recommends -y nodejs
`,
			expectedRst: []ValidateResult{
				{line: 2},
			},
			expectedErr: nil,
		},
		{ // Basic correct case
			dockerfileStr: `FROM ubuntu:nonexistent-hash
RUN apt-get install -y --no-install-recommends python=2.7
`,
			expectedRst: nil,
			expectedErr: nil,
		},
		{ // Options before package installation & contains newlines with \
			dockerfileStr: `FROM ubuntu:nonexistent-hash
RUN \
apt-get update \
&& apt-get install -y --no-install-recommends \
nonexistent-package="2.37-2build1" \
nonpackage="2.42.2-6" \
&& rm -rf /some/path/to/somewhere*
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

		gotRst, gotErr := validateDL3015(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
