package rules

import (
	"testing"
)

func TestDL3023Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian:jesse as build

COPY --from=build some stuff ./
`,
			file: "DL3023Check_Dockerfile",
			expectedRst: []string{
				"DL3023Check_Dockerfile:3 DL3023 COPY --from should reference a previously defined FROM alias\n",
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM debian:jesse as build

RUN stuff

FROM debian:jesse

COPY --from=build some stuff ./
`,
			file:        "DL3023Check_Dockerfile_2",
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3023Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3023Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3023Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3023Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
