package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexDL3033 = regexp.MustCompile(`.+-.+`)

// validateDL3033 Pin versions in yum install. Instead of `yum install <package>` use `yum install <package>-<version>`
func validateDL3033(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isYum, isInstall bool
			l := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "yum":
					isYum = true
				case "install":
					if isYum {
						isInstall = true
					}
				case "&&":
					isYum, isInstall = false, false
					continue
				default:
					if !strings.HasPrefix(v, "-") && isInstall && !regexDL3033.MatchString(v) && l == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
						isYum, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
