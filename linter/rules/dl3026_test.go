package rules

import (
	"testing"
)

func TestValidateDL3026(t *testing.T) {
	cases := []struct {
		name          string
		dockerfileStr string
		opts          *RuleOptions
		expectedRst   []ValidateResult
		expectedErr   error
	}{
		{
			name:          "does not warn on empty allowed registries",
			dockerfileStr: `FROM random.com/debian`,
			opts:          &RuleOptions{TrustedRegistries: []string{}},
			expectedRst:   nil,
			expectedErr:   nil,
		},
		{
			name:          "warn on non-allowed registry",
			dockerfileStr: `FROM random.com/debian:aabbcc`,
			opts:          &RuleOptions{TrustedRegistries: []string{"docker.io"}},
			expectedRst:   []ValidateResult{{line: 1}},
			expectedErr:   nil,
		},
		{
			name:          "does not warn on allowed registries",
			dockerfileStr: `FROM random.com/debian:aabbcc`,
			opts:          &RuleOptions{TrustedRegistries: []string{"x.com", "random.com"}},
			expectedRst:   nil,
			expectedErr:   nil,
		},
		{
			name:          "doesn't warn on scratch image",
			dockerfileStr: `FROM scratch`,
			opts:          &RuleOptions{TrustedRegistries: []string{"x.com", "random.com"}},
			expectedRst:   nil,
			expectedErr:   nil,
		},
		{
			name: "allows all forms of docker.io",
			dockerfileStr: `FROM ubuntu:18.04 AS builder1
FROM zemanlx/ubuntu:18.04 AS builder2
FROM docker.io/zemanlx/ubuntu:18.04 AS builder3`,
			opts:        &RuleOptions{TrustedRegistries: []string{"docker.io"}},
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			name: "allows using previous stages",
			dockerfileStr: `FROM random.com/foo AS builder1
FROM builder1 AS builder2`,
			opts:        &RuleOptions{TrustedRegistries: []string{"random.com"}},
			expectedRst: nil,
			expectedErr: nil,
		},
		{
			name:          "warn on non-allowed wildcard registry",
			dockerfileStr: `FROM x.com/debian`,
			opts:          &RuleOptions{TrustedRegistries: []string{"*.random.com"}},
			expectedRst:   []ValidateResult{{line: 1}},
			expectedErr:   nil,
		},
		{
			name:          "does not warn on allowed wildcard registries",
			dockerfileStr: `FROM foo.random.com/debian`,
			opts:          &RuleOptions{TrustedRegistries: []string{"*.random.com"}},
			expectedRst:   nil,
			expectedErr:   nil,
		},
		{
			name: "does not warn on * registry",
			dockerfileStr: `FROM ubuntu:18.04 AS builder1
FROM zemanlx/ubuntu:18.04 AS builder2
FROM docker.io/zemanlx/ubuntu:18.04 AS builder3`,
			opts:        &RuleOptions{TrustedRegistries: []string{"*"}},
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for _, tc := range cases {
		rst, err := parseDockerfile(tc.dockerfileStr)
		if err != nil {
			t.Errorf("%s parse error %s", tc.name, tc.dockerfileStr)
		}

		gotRst, gotErr := validateDL3026(rst.AST, tc.opts)
		if !isValidateResultEq(gotRst, tc.expectedRst) {
			t.Errorf("%s results deep equal has returned: want %v, got %v", tc.name, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("%s error has returned: want %s, got %s", tc.name, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}
