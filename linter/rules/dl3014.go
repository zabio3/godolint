package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
	"strings"
)

var yesPattern = regexp.MustCompile(`^-[^-]*y.*$`)

// dl3014 Use the `-y` switch to avoid manual input `apt-get -y install <package>`
func dl3014Check(node *parser.Node, file string) (rst []string, err error) {

	for _, child := range node.Children {
		if child.Value == "run" {
			isAptGet, isInstalled := false, false
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
					if isInstalled && !yesPattern.MatchString(v) {
						rst = append(rst, fmt.Sprintf("%s:%v DL3014 Use the `-y` switch to avoid manual input `apt-get -y install <package>`\n", file, child.StartLine))
					}
					isAptGet, isInstalled = false, false
				}
			}
		}
	}
	return rst, nil
}
