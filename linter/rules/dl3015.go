package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3015 Avoid additional packages by specifying --no-install-recommends.
func validateDL3015(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isAptGet, isInstalled, hasRecommends bool
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apt-get":
					isAptGet = true
				case "install":
					if isAptGet {
						isInstalled = true
					}
				case "--no-install-recommends":
					if isInstalled {
						hasRecommends = true
					}
				case "&&":
					if isInstalled && !hasRecommends && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isAptGet, isInstalled = false, false
				default:
					continue
				}
			}

			if isInstalled && !hasRecommends && length == len(rst) {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
