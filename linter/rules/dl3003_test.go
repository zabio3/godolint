package rules

import (
	"testing"
)

func TestDL3003Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM golang:latest

WORKDIR /go
ADD . /go

RUN cd /usr/src/app && git clone git@github.com:zabio3/godolint.git /usr/src/app

CMD ["go", "run", "main.go"]
`,
			file:        "DL3003Check_Dockerfile",
			expectedRst: []string{"DL3003Check_Dockerfile:6 DL3003 Use WORKDIR to switch to a directory\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3003Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3003Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3003Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3003Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
