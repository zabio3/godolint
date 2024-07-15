package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3027 do not use apt; use apt-get or apt-cache instead.
func validateDL3027(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				if v == "apt" && length == len(rst) {
					rst = append(rst, ValidateResult{line: child.StartLine})
				}
			}
		}
	}
	return rst, nil
}
