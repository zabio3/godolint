package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// dl3009 Delete the apt-get lists after installing something.
func dl3009Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			isAptGet, isInstalled, isRm, hasRemove, hasClean := false, false, false, false, false
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apt-get":
					isAptGet = true
				case "install", "update":
					if isAptGet {
						isInstalled = true
					}
				case "clean":
					if isAptGet && isInstalled {
						hasClean = true
					}
				case "&&":
					isAptGet = false
				case "rm":
					if isInstalled {
						isRm = true
					}
				case "/var/lib/apt/lists/*":
					if isRm {
						hasRemove = true
					}
				}
			}
			if isInstalled && !(hasRemove || hasClean) {
				rst = append(rst, fmt.Sprintf("%s:%v DL3009 Delete the apt-get lists after installing something\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
