package rules

import (
	"bytes"
	"testing"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func parseDockerfile(dockerfileStr string) (*parser.Result, error) {
	dockerfile := bytes.NewBufferString(dockerfileStr)
	return parser.Parse(dockerfile)
}

func TestCreateMessage(t *testing.T) {
	cases := []struct {
		rule        *Rule
		vrst        []ValidateResult
		expectedRst []string
	}{
		{
			rule: &Rule{
				Code:         "DL3000",
				Severity:     SeverityError,
				Description:  "Use absolute WORKDIR.",
				ValidateFunc: validateDL3000,
			},
			vrst: []ValidateResult{
				{line: 3},
			},
			expectedRst: []string{
				"#3 DL3000 Use absolute WORKDIR. ",
			},
		},
	}

	for i, tc := range cases {
		gotRst := CreateMessage(tc.rule, tc.vrst)
		if sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d results deep equal has returned: want %v, got %v", i, tc.expectedRst, gotRst)
		}
		cleanup(t)
	}
}

// reflect.DeepEqual(gotRst, expectedRst)
func isValidateResultEq(xs, ys []ValidateResult) bool {
	if (xs == nil) != (ys == nil) {
		return false
	}
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if xs[i] != ys[i] {
			return false
		}
	}
	return true
}

// reflect.DeepEqual(gotRst, gotRst)
func sliceEq(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func cleanup(t *testing.T) {
	t.Helper()
}
