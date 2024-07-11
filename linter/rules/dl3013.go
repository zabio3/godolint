package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexVersion3013 = regexp.MustCompile(`.+[=|@].+`)

// validateDL3013 Pin versions in pip. Instead of `pip install <package>` use `pip install <package>==<version>`.
func validateDL3013(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isPip, isInstall bool
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "pip":
					isPip = true
				case "install":
					if isPip {
						isInstall = true
					}
				case "&&":
					isPip, isInstall = false, false
				default:
					if strings.HasPrefix(v, "--") || strings.HasPrefix(v, "yamllint") {
						continue
					}

					if isPip && isInstall && !regexVersion3013.MatchString(v) && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isPip, isInstall = false, false
				}
			}
		}
	}
	return rst, nil
}
