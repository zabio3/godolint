package rules

import (
	"testing"
)

func TestValidateDL3013(t *testing.T) {
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
			file: "DL3013_Dockerfile",
			expectedRst: []string{
				"DL3013_Dockerfile:2 DL3013 Pin versions in pip. Instead of `pip install <package>` use `pip install <package>==<version>`\n",
				"DL3013_Dockerfile:3 DL3013 Pin versions in pip. Instead of `pip install <package>` use `pip install <package>==<version>`\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM python:3.4
RUN pip install django && pip install https://github.com/Banno/carbon/tarball/0.9.x-fix-events-callback
`,
			file: "DL3013_Dockerfile_2",
			expectedRst: []string{
				"DL3013_Dockerfile_2:2 DL3013 Pin versions in pip. Instead of `pip install <package>` use `pip install <package>==<version>`\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3013(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
