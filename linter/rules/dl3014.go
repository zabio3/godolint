package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var yesPattern = regexp.MustCompile(`^-(y|-yes|-assume-yes)$`)

// validateDL3014 Use the `-y` switch to avoid manual input `apt-get -y install <package>`
func validateDL3014(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isAptGet, isInstalled bool
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apt-get":
					isAptGet = true
				case "install":
					if isAptGet {
						isInstalled = true
					}
				case "&&":
					isAptGet, isInstalled = false, false
				default:
					if isInstalled && !yesPattern.MatchString(v) && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isAptGet, isInstalled = false, false
				}
			}
		}
	}
	return rst, nil
}
