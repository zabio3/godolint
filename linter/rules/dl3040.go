package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3040 Delete the dnf cache after installing something with `dnf clean all`.
func validateDL3040(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN && isDL3040Error(child) {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}

func isDL3040Error(node *parser.Node) bool {
	var isDnf, isInstalled, hasClean bool
	for _, v := range strings.Fields(node.Next.Value) {
		isDnf, isInstalled, hasClean = updateDL3040Status(v, isDnf, isInstalled, hasClean)
	}
	return isInstalled && !hasClean
}

func updateDL3040Status(v string, isDnf, isInstalled, hasClean bool) (bool, bool, bool) {
	switch v {
	case "dnf":
		return true, isInstalled, hasClean
	case "install", "update":
		if isDnf {
			return true, true, hasClean
		}
	case "clean":
		if isDnf && isInstalled {
			return true, true, true
		}
	case "&&":
		if isDnf {
			return false, isInstalled, hasClean
		}
	}

	return isDnf, isInstalled, hasClean
}
