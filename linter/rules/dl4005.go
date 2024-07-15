package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4005 Use SHELL to change the default shell
func validateDL4005(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			isLn := false
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "ln":
					isLn = true
				case "/bin/sh":
					if isLn {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
				}
			}
		}
	}
	return rst, nil
}
