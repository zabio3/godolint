package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// dl3012 Provide an email address or URL as maintainer. (deprecated)
func dl3012Check(node *parser.Node, file string) (rst []string, err error) {
	return rst, nil
}
