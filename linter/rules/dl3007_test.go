package rules

import (
	"testing"
)

func TestDL3007Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian:latest
RUN apt-get update

WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
`,
			file:        "DL3007Check_Dockerfile",
			expectedRst: []string{"DL3007Check_Dockerfile:1 DL3007 Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3007Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3007Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3007Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3007Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
