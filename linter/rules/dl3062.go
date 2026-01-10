package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexVersion3062 = regexp.MustCompile(`.+@.+`)

// validateDL3062 Pin versions in go install. Instead of `go install <package>` use `go install <package>@<version>`
func validateDL3062(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isGo, isInstall bool
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "go":
					isGo = true
				case "install":
					if isGo {
						isInstall = true
					}
				case "&&":
					isGo, isInstall = false, false
					continue
				default:
					if isInstall && !strings.HasPrefix(v, "-") && !regexVersion3062.MatchString(v) && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
						isGo, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
