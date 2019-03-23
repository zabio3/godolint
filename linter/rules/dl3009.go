package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3009 Delete the apt-get lists after installing something.
func validateDL3009(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN && isDL3009Error(child) {
			rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
		}
	}
	return rst, nil
}

func isDL3009Error(node *parser.Node) bool {
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
			if isAptGet {
				isAptGet, hasClean = false, false
			}
		}
	}
	return isInstalled && !(hasRemove || hasClean)
}
