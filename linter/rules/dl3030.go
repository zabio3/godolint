package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3030 Use the `-y` switch to avoid manual input `yum install -y <package>`
func validateDL3030(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isYum, isInstall, hasYesFlag bool
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "yum":
					isYum = true
					hasYesFlag = false
				case "install":
					if isYum {
						isInstall = true
					}
				case "-y":
					if isYum {
						hasYesFlag = true
					}
				case "&&":
					if isYum && isInstall && !hasYesFlag {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isYum, isInstall, hasYesFlag = false, false, false
				}
			}
			if isYum && isInstall && !hasYesFlag {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
