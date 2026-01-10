package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexDL3041 = regexp.MustCompile(`.+-.+`)

// validateDL3041 Pin versions in dnf install. Instead of `dnf install <package>` use `dnf install <package>-<version>`
func validateDL3041(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isDnf, isInstall bool
			l := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "dnf":
					isDnf = true
				case "install":
					if isDnf {
						isInstall = true
					}
				case "&&":
					isDnf, isInstall = false, false
					continue
				default:
					if !strings.HasPrefix(v, "-") && isInstall && !regexDL3041.MatchString(v) && l == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
						isDnf, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
