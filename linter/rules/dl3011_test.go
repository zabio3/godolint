package rules

import (
	"fmt"
	"testing"
)

func TestValidateDL3011(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM debian:1.12.0-stretch

EXPOSE 80000
`, expectedRst: []ValidateResult{
				{line: 3},
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM debian:1.12.0-stretch

EXPOSE hoge
`,
			expectedRst: nil,
			expectedErr: fmt.Errorf("#3 DL3011 not numeric is the value set for the port: hoge"),
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3011(rst.AST, nil)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
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
