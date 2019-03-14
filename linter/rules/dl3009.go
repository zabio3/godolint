package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// dl3009 DL3009 Delete the apt-get lists after installing something.
func dl3009Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" && isDL3009Error(child) {
			rst = append(rst, fmt.Sprintf("%s:%v DL3009 Delete the apt-get lists after installing something\n", file, child.StartLine))
		}
	}
	return rst, nil
}

func isDL3009Error(node *parser.Node) bool {
	if node.Next == nil {
		return false
	}
	isAptGet, isInstalled, isRm, hasRemove, hasClean := false, false, false, false, false
	for _, v := range strings.Fields(node.Next.Value) {
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
		case "rm":
			if isInstalled {
				isRm = true
			}
		case "/var/lib/apt/lists/*":
			if isRm {
				hasRemove = true
			}
		case "&&":
			isAptGet = false
		}
	}
	return isInstalled && !(hasRemove || hasClean)
}
