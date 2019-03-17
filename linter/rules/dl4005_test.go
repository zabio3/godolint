package rules

import (
	"testing"
)

func TestValidateDL4005(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `# Install bash
RUN apk add --update-cache bash=4.3.42-r3

# Use bash as the default shell
RUN ln -sfv /bin/bash /bin/sh
`,
			file: "DL4005_Dockerfile",
			expectedRst: []string{
				"DL4005_Dockerfile:5 DL4005 Use SHELL to change the default shell\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL4005(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
