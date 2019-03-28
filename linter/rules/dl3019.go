package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3019 Use the `--no-cache` switch to avoid the need to use `--update` and remove `/var/cache/apk/*` when done installing packages
func validateDL3019(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			isApk, isUpdate, isRm, hasRemove, length := false, false, false, false, len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apk":
					isApk = true
				case "update":
					if isApk {
						isUpdate = true
					}
				case "rm":
					if isApk {
						isRm = true
					}
				case "/var/cache/apk/*":
					if isRm {
						hasRemove = true
					}
				case "&&":
					isApk = false
					continue
				default:
					if (isUpdate || hasRemove) && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
						isApk = false
					}
				}
			}
		}
	}
	return rst, nil
}
