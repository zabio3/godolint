package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3012 Provide an email address or URL as maintainer. (deprecated)
func validateDL3012(node *parser.Node, file string) (rst []string, err error) {
	return rst, nil
}
