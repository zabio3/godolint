package rules

import (
	"testing"
)

func TestDL4006Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `RUN wget -O - https://some.site | wc -l > /number
`,
			file: "DL4006Check_Dockerfile",
			expectedRst: []string{
				"DL4006Check_Dockerfile:1 DL4006 Set the `SHELL` option -o pipefail before `RUN` with a pipe in it\n",
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN wget -O - https://some.site | wc -l > /number
`,
			file:        "DL4006Check_Dockerfile_2",
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl4006Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
