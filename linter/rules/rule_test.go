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

func cleanup(t *testing.T) {
	t.Helper()
}
