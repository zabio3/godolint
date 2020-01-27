package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3017 Do not use apk upgrade
func validateDL3017(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isApk bool
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apk":
					isApk = true
				case "upgrade":
					if isApk && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
				case "&&":
					isApk = false
					continue
				default:
				}
			}
		}
	}
	return rst, nil
}
