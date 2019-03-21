package rules

import (
	"fmt"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3005 is "Do not use apt-get upgrade or dist-upgrade."
func validateDL3005(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			isAptGet, isUpgrade := false, false
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apt-get":
					isAptGet = true
				case "upgrade":
					isUpgrade = true
				}
			}
			if isAptGet && isUpgrade {
				rst = append(rst, fmt.Sprintf("%s:%v DL3005 Do not use apt-get upgrade or dist-upgrade.\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
