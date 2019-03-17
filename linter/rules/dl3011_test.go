package rules

import (
	"fmt"
	"testing"
)

func TestValidateDL3011(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian:1.12.0-stretch

EXPOSE 80000
`,
			file:        "DL3011_Dockerfile",
			expectedRst: []string{"DL3011_Dockerfile:3 DL3011 Valid UNIX ports range from 0 to 65535\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM debian:1.12.0-stretch

EXPOSE hoge
`,
			file:        "DL3011_Dockerfile_2",
			expectedRst: nil,
			expectedErr: fmt.Errorf("DL3011_Dockerfile_2:3 DL3011 not numeric is the value set for the port: hoge"),
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3011(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		expectedErr := tc.expectedErr
		if gotErr != nil && expectedErr != nil {
			if gotErr.Error() != expectedErr.Error() {
				t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
			}
		} else {
			if gotErr != tc.expectedErr {
				t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
			}
		}

		cleanup(t)
	}
}
