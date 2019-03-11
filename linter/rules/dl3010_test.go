package rules

import (
	"testing"
)

func TestDL3010Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian:1.12.0-stretch

COPY hoge.tar.xz /

CMD ["go", "run", "main.go"]
`,
			file:        "DL3010Check_Dockerfile",
			expectedRst: []string{"DL3010Check_Dockerfile:3 DL3010 Use ADD for extracting archives into an image.\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3010Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3010Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3010Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3010Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
