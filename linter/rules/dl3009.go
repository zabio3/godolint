package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3009 Delete the apt-get lists after installing something.
func validateDL3009(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN && isDL3009Error(child) {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}

func isDL3009Error(node *parser.Node) bool {
	var isAptGet, isInstalled, hasClean, isRm, hasRemove bool
	for _, v := range strings.Fields(node.Next.Value) {
		isAptGet, isInstalled, hasClean, isRm, hasRemove = updateDL3009Status(v, isAptGet, isInstalled, hasClean, isRm, hasRemove)
	}
	return isInstalled && !(hasRemove || hasClean)
}

func updateDL3009Status(v string, isAptGet, isInstalled, hasClean, isRm, hasRemove bool) (bool, bool, bool, bool, bool) {
	switch v {
	case "apt-get":
		return true, isInstalled, hasClean, isRm, hasRemove
	case "install", "update":
		if isAptGet {
			return true, true, hasClean, isRm, hasRemove
		}
	case "clean":
		if isAptGet && isInstalled {
			return true, true, true, isRm, hasRemove
		}
	case "rm":
		if isInstalled {
			return true, true, hasClean, hasRemove, true
		}
	case "/var/lib/apt/lists/*":
		if isRm {
			return true, true, hasClean, true, true
		}
	case "&&":
		if isAptGet {
			return false, isInstalled, false, isRm, hasRemove
		}
	}

	return isAptGet, isInstalled, hasClean, isRm, hasRemove
}
