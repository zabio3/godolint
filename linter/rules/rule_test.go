package rules

import (
	"bytes"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func dockerFileParse(dockerfileStr string) (*parser.Result, error) {
	dockerfile := bytes.NewBufferString(dockerfileStr)
	return parser.Parse(dockerfile)
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
