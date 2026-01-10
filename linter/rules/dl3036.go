package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3036 Delete the zypper cache after installing something.
func validateDL3036(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN && isDL3036Error(child) {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}

func isDL3036Error(node *parser.Node) bool {
	var isZypper, isInstalled, hasClean bool
	for _, v := range strings.Fields(node.Next.Value) {
		isZypper, isInstalled, hasClean = updateDL3036Status(v, isZypper, isInstalled, hasClean)
	}
	return isInstalled && !hasClean
}

func updateDL3036Status(v string, isZypper, isInstalled, hasClean bool) (bool, bool, bool) {
	switch v {
	case "zypper":
		return true, isInstalled, hasClean
	case "install", "update", "in", "up":
		if isZypper {
			return true, true, hasClean
		}
	case "clean", "cc":
		if isZypper && isInstalled {
			return true, true, true
		}
	case "&&":
		if isZypper {
			return false, isInstalled, hasClean
		}
	}

	return isZypper, isInstalled, hasClean
}
