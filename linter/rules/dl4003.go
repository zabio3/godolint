package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4003 Multiple CMD instructions found.
func validateDL4003(node *parser.Node) (rst []ValidateResult, err error) {
	isCmd := false
	for _, child := range node.Children {
		if child.Value == CMD {
			if !isCmd {
				isCmd = true
			} else {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
