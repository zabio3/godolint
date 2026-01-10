package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexVersion3028 = regexp.MustCompile(`.+:.+`)

// validateDL3028 Pin versions in gem install. Instead of `gem install <package>` use `gem install <package>:<version>` or `gem install <package> -v <version>`
func validateDL3028(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			fields := strings.Fields(child.Next.Value)
			var isGem, isInstall, hasVersionFlag bool
			var packageName string
			length := len(rst)

			for i := 0; i < len(fields); i++ {
				v := fields[i]
				switch v {
				case "gem":
					isGem = true
					hasVersionFlag = false
					packageName = ""
				case "install":
					if isGem {
						isInstall = true
					}
				case "&&":
					if isGem && isInstall && packageName != "" && !hasVersionFlag && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isGem, isInstall, hasVersionFlag = false, false, false
					packageName = ""
				case "-v", "--version":
					if isInstall {
						hasVersionFlag = true
					}
				default:
					if isInstall && !strings.HasPrefix(v, "-") {
						if regexVersion3028.MatchString(v) {
							// Package has version in format package:version
							hasVersionFlag = true
						} else {
							// Package without version
							packageName = v
						}
					}
				}
			}
			// Check at end of RUN command
			if isGem && isInstall && packageName != "" && !hasVersionFlag && length == len(rst) {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
