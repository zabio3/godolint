package rules

import (
	"testing"
)

func TestValidateDL3022(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian:jesse

RUN stuff

FROM debian:jesse

COPY --from=build some stuff ./
`,
			file: "DL3022_Dockerfile",
			expectedRst: []string{
				"DL3022_Dockerfile:5 DL3022 COPY --from should reference a previously defined FROM alias\n",
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM debian:jesse as build

RUN stuff

FROM debian:jesse

COPY --from=build some stuff ./
`,
			file:        "DL3022_Dockerfile",
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3022(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
