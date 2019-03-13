package rules

import (
	"testing"
)

func TestDL3018Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM alpine:3.7
RUN apk --no-cache add foo
`,
			file: "DL3018Check_Dockerfile",
			expectedRst: []string{
				"DL3018Check_Dockerfile:2 DL3018 Pin versions in apk add. Instead of `apk add <package>` use `apk add <package>=<version>`\n",
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM alpine:3.7
RUN apk --no-cache add foo && bar
`,
			file: "DL3018Check_Dockerfile",
			expectedRst: []string{
				"DL3018Check_Dockerfile:2 DL3018 Pin versions in apk add. Instead of `apk add <package>` use `apk add <package>=<version>`\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3018Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3018Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3018Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3018Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
