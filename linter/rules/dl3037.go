package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexDL3037 = regexp.MustCompile(`.+=.+`)

// validateDL3037 Pin versions in zypper install. Instead of `zypper install <package>` use `zypper install <package>=<version>`
func validateDL3037(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isZypper, isInstall bool
			l := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "zypper":
					isZypper = true
				case "install", "in":
					if isZypper {
						isInstall = true
					}
				case "&&":
					isZypper, isInstall = false, false
					continue
				default:
					if !strings.HasPrefix(v, "-") && isInstall && !regexDL3037.MatchString(v) && l == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
						isZypper, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
