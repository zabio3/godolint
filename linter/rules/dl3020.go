package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3020 Use COPY instead of ADD for files and folders
func validateDL3020(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == ADD {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}
