package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3019 Use the `--no-cache` switch to avoid the need to use `--update` and remove `/var/cache/apk/*` when done installing packages
func validateDL3019(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			isApk, isAdd, length := false, false, len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apk":
					isApk = true
				case "add":
					if isApk {
						isAdd = true
					}
				case "&&":
					isApk, isAdd = false, false
					continue
				default:
					if isAdd && v != "--update" && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
						isApk, isAdd = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
