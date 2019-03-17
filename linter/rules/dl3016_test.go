package rules

import (
	"testing"
)

func TestValidateDL3016(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
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
`,
			file: "DL3016_Dockerfile",
			expectedRst: []string{
				"DL3016_Dockerfile:3 DL3016 Pin versions in npm. Instead of `npm install <package>` use `npm install <package>@<version>`\n",
				"DL3016_Dockerfile:4 DL3016 Pin versions in npm. Instead of `npm install <package>` use `npm install <package>@<version>`\n",
				"DL3016_Dockerfile:5 DL3016 Pin versions in npm. Instead of `npm install <package>` use `npm install <package>@<version>`\n",
				"DL3016_Dockerfile:6 DL3016 Pin versions in npm. Instead of `npm install <package>` use `npm install <package>@<version>`\n",
				"DL3016_Dockerfile:7 DL3016 Pin versions in npm. Instead of `npm install <package>` use `npm install <package>@<version>`\n",
				"DL3016_Dockerfile:8 DL3016 Pin versions in npm. Instead of `npm install <package>` use `npm install <package>@<version>`\n",
				"DL3016_Dockerfile:9 DL3016 Pin versions in npm. Instead of `npm install <package>` use `npm install <package>@<version>`\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3016(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
