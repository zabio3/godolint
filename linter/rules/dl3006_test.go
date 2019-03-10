package rules

import (
	"testing"
)

func TestDL3006Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian
RUN apt-get update

WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
`,
			file:        "DL3006Check_Dockerfile",
			expectedRst: []string{"DL3006Check_Dockerfile:1 DL3006 Always tag the version of an image explicitly\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3006Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3006Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3006Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3006Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
