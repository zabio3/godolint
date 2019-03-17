package rules

import (
	"testing"
)

func TestValidateDL3021(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM node:carbon
COPY package.json yarn.lock my_app
`,
			file: "DL3021_Dockerfile",
			expectedRst: []string{
				"DL3021_Dockerfile:2 DL3021 `COPY` with more than 2 arguments requires the last argument to end with `/`\n",
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM node:carbon
COPY package.json yarn.lock my_app/
`,
			file:        "DL3021_Dockerfile_2",
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3021(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
