package rules

import (
	"testing"
)

func TestValidateDL3016(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM node:8.9.1

RUN npm install express
RUN npm install @myorg/privatepackage
RUN npm install express sax@0.1.1
RUN npm install --global express
RUN npm install git+ssh://git@github.com:npm/npm.git
RUN npm install git+http://isaacs@github.com/npm/npm && npm install git+https://isaacs@github.com/npm/npm.git
RUN npm install git://github.com/npm/npm.git
`, expectedRst: []ValidateResult{
				{line: 3, addMsg: ""},
				{line: 4, addMsg: ""},
				{line: 5, addMsg: ""},
				{line: 6, addMsg: ""},
				{line: 7, addMsg: ""},
				{line: 8, addMsg: ""},
				{line: 9, addMsg: ""},
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3016(rst.AST)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
