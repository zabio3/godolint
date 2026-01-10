package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3034 Use the `-n` switch to avoid manual input `zypper -n install <package>`
func validateDL3034(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isZypper, isInstall, hasNonInteractive bool
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "zypper":
					isZypper = true
					hasNonInteractive = false
				case "install", "update", "remove", "in", "up", "rm":
					if isZypper {
						isInstall = true
					}
				case "-n", "--non-interactive":
					if isZypper {
						hasNonInteractive = true
					}
				case "&&":
					if isZypper && isInstall && !hasNonInteractive {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isZypper, isInstall, hasNonInteractive = false, false, false
				}
			}
			if isZypper && isInstall && !hasNonInteractive {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
