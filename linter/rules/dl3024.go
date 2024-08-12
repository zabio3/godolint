package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3024 FROM aliases (stage names) must be unique
func validateDL3024(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	var isAs bool
	var asBuildName []string
	for _, child := range node.Children {
		if child.Value == FROM {
			for _, v := range strings.Fields(child.Original) {
				switch v {
				case "as":
					isAs = true
				default:
					if isAs {
						if isContain(asBuildName, v) {
							rst = append(rst, ValidateResult{line: child.StartLine})
						} else {
							asBuildName = append(asBuildName, v)
						}
						isAs = false
					}
				}
			}
		}
	}
	return rst, nil
}
