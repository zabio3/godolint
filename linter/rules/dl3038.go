package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3038 Use the `-y` switch to avoid manual input `dnf install -y <package>`
func validateDL3038(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isDnf, isInstall, hasYesFlag bool
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "dnf":
					isDnf = true
					hasYesFlag = false
				case "install":
					if isDnf {
						isInstall = true
					}
				case "-y":
					if isDnf {
						hasYesFlag = true
					}
				case "&&":
					if isDnf && isInstall && !hasYesFlag {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isDnf, isInstall, hasYesFlag = false, false, false
				}
			}
			if isDnf && isInstall && !hasYesFlag {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
