package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3032 Delete the yum cache after installing something with `yum clean all`.
func validateDL3032(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN && isDL3032Error(child) {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}

func isDL3032Error(node *parser.Node) bool {
	var isYum, isInstalled, hasClean bool
	for _, v := range strings.Fields(node.Next.Value) {
		isYum, isInstalled, hasClean = updateDL3032Status(v, isYum, isInstalled, hasClean)
	}
	return isInstalled && !hasClean
}

func updateDL3032Status(v string, isYum, isInstalled, hasClean bool) (bool, bool, bool) {
	switch v {
	case "yum":
		return true, isInstalled, hasClean
	case "install", "update":
		if isYum {
			return true, true, hasClean
		}
	case "clean":
		if isYum && isInstalled {
			return true, true, true
		}
	case "&&":
		if isYum {
			return false, isInstalled, hasClean
		}
	}

	return isYum, isInstalled, hasClean
}
