package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
	"strings"
)

var verPattern3008 = regexp.MustCompile(`.+=.+`)

// dl3008 Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`
func dl3008Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			isAptGet, isInstall := false, false
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
					if isInstall && !verPattern3008.MatchString(v) {
						rst = append(rst, fmt.Sprintf("%s:%v DL3008 Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`\n", file, child.StartLine))
						isAptGet, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
