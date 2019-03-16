package rules

import (
	"testing"
)

func TestDL4000Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM busybox
MAINTAINER zabio3 <zabio1192@gmail.com>
`,
			file: "DL4000Check_Dockerfile",
			expectedRst: []string{
				"DL4000Check_Dockerfile:2 DL4000 MAINTAINER is deprecated\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl4000Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl4000Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl4000Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl4000Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
