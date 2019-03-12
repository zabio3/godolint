package rules

import (
	"testing"
)

func TestDL3014Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian
RUN apt-get install python=2.7
`,
			file:        "DL3014Check_Dockerfile",
			expectedRst: []string{"DL3014Check_Dockerfile:2 DL3014 Use the `-y` switch to avoid manual input `apt-get -y install <package>`\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM debian
RUN apt-get install python=2.7 && apt-get install ruby
`,
			file:        "DL3014Check_Dockerfile_2",
			expectedRst: []string{"DL3014Check_Dockerfile_2:2 DL3014 Use the `-y` switch to avoid manual input `apt-get -y install <package>`\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3014Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3014Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3014Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3014Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
