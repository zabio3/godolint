package rules

import (
	"testing"
)

func TestDL4001Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian
RUN wget http://google.com
RUN curl http://bing.com
`,
			file: "DL4001Check_Dockerfile",
			expectedRst: []string{
				"DL4001Check_Dockerfile:3 DL4001 Either use Wget or Curl but not both\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl4001Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl4001Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl4001Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl4001Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
