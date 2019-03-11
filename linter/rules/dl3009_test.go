package rules

import (
	"testing"
)

func TestDL3009Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian:1.12.0-stretch
RUN apt-get update && apt-get install -y python

CMD ["go", "run", "main.go"]
`,
			file:        "DL3009Check_Dockerfile",
			expectedRst: []string{"DL3009Check_Dockerfile:2 DL3009 Delete the apt-get lists after installing something\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM debian:1.12.0-stretch
RUN apt-get update && apt-get install -y python && apt-get clean && rm /var/lib/apt/lists/*

CMD ["go", "run", "main.go"]
`,
			file:        "DL3009Check_Dockerfile_2",
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3009Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3009Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3009Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3009Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
