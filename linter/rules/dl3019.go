package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3019 Use the `--no-cache` switch to avoid the need to use `--update` and remove `/var/cache/apk/*` when done installing packages
func validateDL3019(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN && isDL3019Error(child) {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}

func isDL3019Error(node *parser.Node) bool {
	var isApk, isRm bool
	for _, v := range strings.Fields(node.Next.Value) {
		switch v {
		case "apk":
			isApk = true
		case "update":
			if isApk {
				return true
			}
		case "rm":
			if isApk {
				isRm = true
			}
		case "/var/cache/apk/*":
			if isRm {
				return true
			}
		case "&&":
			isApk = false
		}
	}

	return false
}
