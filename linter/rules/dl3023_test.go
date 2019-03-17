package rules

import (
	"testing"
)

func TestValidateDL3023(t *testing.T) {
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
			file: "DL3023_Dockerfile",
			expectedRst: []string{
				"DL3023_Dockerfile:3 DL3023 COPY --from should reference a previously defined FROM alias\n",
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM debian:jesse as build

RUN stuff

FROM debian:jesse

COPY --from=build some stuff ./
`,
			file:        "DL3023_Dockerfile_2",
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3023(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
