package rules

import (
	"testing"
)

func TestDL3021Check(t *testing.T) {
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
			file: "DL3021Check_Dockerfile",
			expectedRst: []string{
				"DL3021Check_Dockerfile:2 DL3021 `COPY` with more than 2 arguments requires the last argument to end with `/`\n",
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM node:carbon
COPY package.json yarn.lock my_app/
`,
			file:        "DL3021Check_Dockerfile_2",
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3021Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3021Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3021Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3021Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
