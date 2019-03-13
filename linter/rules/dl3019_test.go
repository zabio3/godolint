package rules

import (
	"testing"
)

func TestDL3019Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM alpine:3.7
RUN apk update \
    && apk add foo=1.0 \
    && rm -rf /var/cache/apk/*
`,
			file: "DL3019Check_Dockerfile",
			expectedRst: []string{
				"DL3019Check_Dockerfile:2 DL3019 Use the `--no-cache` switch to avoid the need to use `--update` and remove `/var/cache/apk/*` when done installing packages\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3019Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3019Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3019Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3019Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
