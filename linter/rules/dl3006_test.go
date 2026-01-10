package rules

import (
	"testing"
)

func TestValidateDL3006(t *testing.T) {
	cases := []struct {
		name          string
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			name: "warns on image without tag",
			dockerfileStr: `FROM debian
RUN apt-get update

WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
`,
			expectedRst: []ValidateResult{
				{line: 1},
			},
			expectedErr: nil,
		},
		{
			name:          "doesn't warn on scratch image",
			dockerfileStr: `FROM scratch`,
			expectedRst:   nil,
			expectedErr:   nil,
		},
		{
			name: "doesn't warn on image with tag",
			dockerfileStr: `FROM debian:11
RUN apt-get update
`,
			expectedRst:   nil,
			expectedErr:   nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rst, err := parseDockerfile(tc.dockerfileStr)
			if err != nil {
				t.Errorf("parse error %s", tc.dockerfileStr)
			}

			gotRst, gotErr := validateDL3006(rst.AST, nil)
			if !isValidateResultEq(gotRst, tc.expectedRst) {
				t.Errorf("results deep equal has returned: want %v, got %v", tc.expectedRst, gotRst)
			}

			if gotErr != tc.expectedErr {
				t.Errorf("error has returned: want %s, got %s", tc.expectedErr, gotErr)
			}
			cleanup(t)
		})
	}
}
