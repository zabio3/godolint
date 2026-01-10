package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3061 Invalid instruction order. Dockerfile must begin with `FROM`, `ARG` or comment.
func validateDL3061(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	if len(node.Children) == 0 {
		return rst, nil
	}

	firstChild := node.Children[0]
	if firstChild.Value != FROM && firstChild.Value != ARG {
		rst = append(rst, ValidateResult{line: firstChild.StartLine})
	}

	return rst, nil
}
