package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3003 is "Use WORKDIR to switch to a directory"
func validateDL3003(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			for _, v := range strings.Fields(child.Next.Value) {
				if v == "cd" {
					rst = append(rst, ValidateResult{line: child.StartLine})
				}
				break
			}
		}
	}
	return rst, nil
}
