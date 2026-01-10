package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3004 validates "Do not use sudo as it leads to unpredictable behavior".
// It checks if the first command in a RUN instruction is "sudo".
func validateDL3004(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			fields := strings.Fields(child.Next.Value)
			if len(fields) > 0 && fields[0] == "sudo" {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
