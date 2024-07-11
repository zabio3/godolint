package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexVersion3018 = regexp.MustCompile(`.+=.+`)

// validateDL3018 Do not use apk upgrade
func validateDL3018(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isApk, isAdd bool
			length := len(rst)
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
				default:
					if isAdd && !regexVersion3018.MatchString(v) && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
						isApk, isAdd = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
