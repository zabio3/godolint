package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
	"strings"
)

var verPattern = regexp.MustCompile(`.+=.+`)

// dl3008 Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`
func dl3008Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			isAptGet, isInstall := false, false
			args := strings.Fields(child.Next.Value)
			for _, v := range args {
				switch v {
				case "apt-get":
					isAptGet = true
				case "install":
					if isAptGet {
						isInstall = true
					}
				case "&&":
					isAptGet, isInstall = false, false
				default:
					if isInstall && !verPattern.MatchString(v) {
						rst = append(rst, fmt.Sprintf("%s:%v DL3008 Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`\n", file, child.StartLine))
						isAptGet, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}