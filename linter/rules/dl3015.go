package rules

import (
	"fmt"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3015 Avoid additional packages by specifying --no-install-recommends.
func validateDL3015(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			isAptGet, isInstalled, length := false, false, len(rst)
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
					if isInstalled && v != "--no-install-recommends" && length == len(rst) {
						rst = append(rst, fmt.Sprintf("%s:%v DL3015 Avoid additional packages by specifying `--no-install-recommends`\n", file, child.StartLine))
					}
					isAptGet, isInstalled = false, false
				}
			}
		}
	}
	return rst, nil
}
