package rules

import (
	"testing"
)

func TestDL3002Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM golang:latest

USER root
WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
`,
			file:        "DL3002Check_Dockerfile",
			expectedRst: []string{"DL3002Check_Dockerfile:3 DL3002 Last USER should not be root\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM golang:latest

USER root
WORKDIR /go
ADD . /go
USER zabio3

CMD ["go", "run", "main.go"]
`,
			file:        "DL3002Check_Dockerfile_2",
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3002Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3002Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3002Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3002Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
