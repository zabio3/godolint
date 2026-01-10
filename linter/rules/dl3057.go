package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3057 `HEALTHCHECK` instruction missing.
func validateDL3057(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	hasHealthcheck := false

	for _, child := range node.Children {
		if child.Value == HEALTHCHECK {
			hasHealthcheck = true
			break
		}
	}

	if !hasHealthcheck {
		rst = append(rst, ValidateResult{line: 1, addMsg: "No HEALTHCHECK instruction found"})
	}

	return rst, nil
}
