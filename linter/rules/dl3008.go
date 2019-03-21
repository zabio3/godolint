package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexDL3008 = regexp.MustCompile(`.+=.+`)

// validateDL3008 Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`
func validateDL3008(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			isAptGet, isInstall, length := false, false, len(rst)
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
					if isInstall && !regexDL3008.MatchString(v) && length == len(rst) {
						rst = append(rst, fmt.Sprintf("%s:%v DL3008 Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`\n", file, child.StartLine))
						isAptGet, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
