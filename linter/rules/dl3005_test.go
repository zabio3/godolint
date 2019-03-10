package rules

import (
	"testing"
)

func TestDL3005Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian
RUN apt-get update && apt-get upgrade

WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
`,
			file:        "DL3005Check_Dockerfile",
			expectedRst: []string{"DL3005Check_Dockerfile:2 DL3005 Do not use apt-get upgrade or dist-upgrade.\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3005Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3005Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3005Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3005Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
