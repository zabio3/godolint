package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3060 Delete the yarn cache after installing something.
func validateDL3060(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN && isDL3060Error(child) {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}

func isDL3060Error(node *parser.Node) bool {
	var isYarn, isInstalled, hasClean bool
	fields := strings.Fields(node.Next.Value)
	for i := 0; i < len(fields); i++ {
		v := fields[i]
		switch v {
		case "yarn":
			isYarn = true
		case "install", "add":
			if isYarn {
				isInstalled = true
			}
			isYarn = false
		case "cache":
			if isYarn && i+1 < len(fields) && fields[i+1] == "clean" {
				hasClean = true
			}
		case "&&":
			isYarn = false
		default:
			if isYarn && v != "cache" {
				isYarn = false
			}
		}
	}
	return isInstalled && !hasClean
}
