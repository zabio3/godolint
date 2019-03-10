package rules

import (
	"testing"
)

func TestDL3000Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM golang:latest

WORKDIR go/
ADD . /go

CMD ["go", "run", "main.go"]
`,
			file:        "DL3000Check_Dockerfile",
			expectedRst: []string{"DL3000Check_Dockerfile:3 DL3000 Use absolute WORKDIR\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3000Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3000Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3000Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3000Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
