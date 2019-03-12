package rules

import (
	"testing"
)

func TestDL3017Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM alpine:3.7
RUN apk update \
    && apk upgrade \
    && apk add foo=1.0 \
    && rm -rf /var/cache/apk/*
`,
			file: "DL3017Check_Dockerfile",
			expectedRst: []string{
				"DL3017Check_Dockerfile:2 3017 Do not use apk upgrade\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3017Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3017Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3017Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3017Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
