package rules

import (
	"fmt"
	"testing"
)

func TestDL3011Check(t *testing.T) {
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
			file:        "DL3011Check_Dockerfile",
			expectedRst: []string{"DL3011Check_Dockerfile:3 DL3011 Valid UNIX ports range from 0 to 65535\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM debian:1.12.0-stretch

EXPOSE hoge
`,
			file:        "DL3011Check_Dockerfile_2",
			expectedRst: nil,
			expectedErr: fmt.Errorf("DL3011Check_Dockerfile_2:3 DL3011 not numeric is the value set for the port: hoge"),
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3011Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3011Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3011Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		expectedErr := tc.expectedErr
		if gotErr != nil && expectedErr != nil {
			if gotErr.Error() != expectedErr.Error() {
				t.Errorf("#%d dl3011Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
			}
		} else {
			if gotErr != tc.expectedErr {
				t.Errorf("#%d dl3011Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
			}
		}

		cleanup(t)
	}
}
