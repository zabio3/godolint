package rules

import (
	"testing"
)

func TestValidateDL3008(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM busybox:1.0
RUN apt-get install python

CMD ["go", "run", "main.go"]
`,
			file:        "DL3008_Dockerfile",
			expectedRst: []string{"DL3008_Dockerfile:2 DL3008 Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM busybox:1.0
RUN apt-get install python && apt-get clean

CMD ["go", "run", "main.go"]
`,
			file:        "DL3008_Dockerfile_2",
			expectedRst: []string{"DL3008_Dockerfile_2:2 DL3008 Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3008(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
