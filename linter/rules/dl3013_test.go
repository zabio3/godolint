package rules

import (
	"testing"
)

func TestDL3013Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM python:3.4
RUN pip install django
RUN pip install https://github.com/Banno/carbon/tarball/0.9.x-fix-events-callback
`,
			file: "DL3013Check_Dockerfile",
			expectedRst: []string{
				"DL3013Check_Dockerfile:2 Pin versions in pip. Instead of `pip install <package>` use `pip install <package>==<version>`\n",
				"DL3013Check_Dockerfile:3 Pin versions in pip. Instead of `pip install <package>` use `pip install <package>==<version>`\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM python:3.4
RUN pip install django && pip install https://github.com/Banno/carbon/tarball/0.9.x-fix-events-callback
`,
			file: "DL3013Check_Dockerfile_2",
			expectedRst: []string{
				"DL3013Check_Dockerfile_2:2 Pin versions in pip. Instead of `pip install <package>` use `pip install <package>==<version>`\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3013Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3013Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3013Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3013Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
