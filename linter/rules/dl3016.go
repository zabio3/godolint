package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexVersion3016 = regexp.MustCompile(`.+[#|@][0-9"]+`)

// validateDL3016 Pin versions in npm. Instead of `npm install <package>` use `npm install <package>@<version>`
func validateDL3016(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isNpm, isInstall bool
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "npm":
					isNpm = true
				case "install":
					if isNpm {
						isInstall = true
					}
				case "&&":
					isNpm, isInstall = false, false
					continue
				default:
					if isInstall && !regexVersion3016.MatchString(v) && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
						isNpm, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
