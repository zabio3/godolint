package rules

import (
	"testing"
)

func TestDL3024Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian:jesse as build

RUN stuff

FROM debian:jesse as build

RUN more_stuff
`,
			file: "DL3024Check_Dockerfile",
			expectedRst: []string{
				"DL3024Check_Dockerfile:5 DL3024 FROM aliases (stage names) must be unique\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3024Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3024Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3024Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3024Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
