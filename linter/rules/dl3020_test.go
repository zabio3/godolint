package rules

import (
	"testing"
)

func TestDL3020Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM python:3.4
ADD requirements.txt /usr/src/app/
`,
			file: "DL3020Check_Dockerfile",
			expectedRst: []string{
				"DL3020Check_Dockerfile:2 DL3020 Use COPY instead of ADD for files and folders\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3020Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3020Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3020Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3020Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
