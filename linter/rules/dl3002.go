package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3002 is "Last user should not be root."
func validateDL3002(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	var isLastRootUser bool
	var lastRootUserPos int
	for _, child := range node.Children {
		if child.Value == USER {
			if child.Next.Value == "root" || child.Next.Value == "0" {
				isLastRootUser = true
				lastRootUserPos = child.StartLine
				continue
			}
			isLastRootUser = false
		}
	}
	if isLastRootUser {
		rst = append(rst, ValidateResult{line: lastRootUserPos})
	}
	return rst, nil
}
