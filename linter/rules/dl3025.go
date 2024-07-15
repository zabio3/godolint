package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3025 Use arguments JSON notation for CMD and ENTRYPOINT arguments
func validateDL3025(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if (child.Value == ENTRYPOINT) || (child.Value == CMD) {
			l := len(child.Value)
			if child.Original[l+1:l+2] != "[" ||
				child.Original[len(child.Original)-1:] != "]" {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
