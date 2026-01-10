package rules

import (
	"slices"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3024 validates "FROM aliases (stage names) must be unique".
func validateDL3024(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	var nextIsAlias bool
	var aliases []string
	for _, child := range node.Children {
		if child.Value != FROM {
			continue
		}
		for _, token := range strings.Fields(child.Original) {
			if strings.EqualFold(token, "as") {
				nextIsAlias = true
				continue
			}
			if nextIsAlias {
				if slices.Contains(aliases, token) {
					rst = append(rst, ValidateResult{line: child.StartLine})
				} else {
					aliases = append(aliases, token)
				}
				nextIsAlias = false
			}
		}
	}
	return rst, nil
}
