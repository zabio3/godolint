package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3004 is "Do not use sudo as it leads to unpredictable behavior. Use a tool like gosu to enforce root."
func validateDL3004(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			for _, v := range strings.Fields(child.Next.Value) {
				if v == "sudo" {
					rst = append(rst, ValidateResult{line: child.StartLine})
				}
				break
			}
		}
	}
	return rst, nil
}
