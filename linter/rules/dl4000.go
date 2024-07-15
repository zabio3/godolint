package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4000 MAINTAINER is deprecated
func validateDL4000(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == MAINTAINER {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}
