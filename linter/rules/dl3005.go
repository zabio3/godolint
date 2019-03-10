package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// dl3005Check is "Do not use apt-get upgrade or dist-upgrade."
func dl3005Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			isAptGet, isUpgrade := false, false
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apt-get":
					isAptGet = true
				case "upgrade":
					isUpgrade = true
				default:
				}
			}
			if isAptGet && isUpgrade {
				rst = append(rst, fmt.Sprintf("%s:%v DL3005 Do not use apt-get upgrade or dist-upgrade.\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
