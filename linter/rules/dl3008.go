package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexDL3008 = regexp.MustCompile(`.+=.+`)

// validateDL3008 Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`
func validateDL3008(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isAptGet, isInstall bool
			l := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apt-get":
					isAptGet = true
				case "install":
					if isAptGet {
						isInstall = true
					}
				case "&&":
					isAptGet, isInstall = false, false
					continue
				default:
					if !strings.HasPrefix(v, "-") && isInstall && !regexDL3008.MatchString(v) && l == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
						isAptGet, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
