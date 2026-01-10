package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3003 validates "Use WORKDIR to switch to a directory".
// It checks if the first command in a RUN instruction is "cd".
func validateDL3003(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			fields := strings.Fields(child.Next.Value)
			if len(fields) > 0 && fields[0] == "cd" {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
