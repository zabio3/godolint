package rules

import (
	"testing"
)

func TestValidateDL3004(t *testing.T) {
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
			file:        "DL3004_Dockerfile",
			expectedRst: []string{"DL3004_Dockerfile:3 DL3004 Do not use sudo as it leads to unpredictable behavior. Use a tool like gosu to enforce root.\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3004(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
