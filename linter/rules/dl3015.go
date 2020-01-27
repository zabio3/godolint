package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3015 Avoid additional packages by specifying --no-install-recommends.
func validateDL3015(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isAptGet, isInstalled bool
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apt-get":
					isAptGet = true
				case "install":
					if isAptGet {
						isInstalled = true
					}
				case "&&":
					isAptGet, isInstalled = false, false
				default:
					if isInstalled && v != "--no-install-recommends" && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isAptGet, isInstalled = false, false
				}
			}
		}
	}
	return rst, nil
}
