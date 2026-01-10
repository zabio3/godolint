package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3035 is "Do not use zypper dist-upgrade."
func validateDL3035(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isZypper, isDistUpgrade bool
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "zypper":
					isZypper = true
				case "dist-upgrade", "dup":
					isDistUpgrade = true
				}
			}
			if isZypper && isDistUpgrade {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
