package rules

import (
	"testing"
)

func TestDL3004Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM ubuntu:latest

RUN sudo apt-get install

WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
`,
			file:        "DL3004Check_Dockerfile",
			expectedRst: []string{"DL3004Check_Dockerfile:3 DL3004 Do not use sudo as it leads to unpredictable behavior. Use a tool like gosu to enforce root.\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3004Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3004Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3004Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3004Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
