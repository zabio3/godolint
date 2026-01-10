package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3049 Label `<label>` is missing.
// This is a basic implementation that checks if any labels exist.
func validateDL3049(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	foundLabels := make(map[string]bool)

	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)
			for key := range pairs {
				foundLabels[key] = true
			}
		}
	}

	if len(foundLabels) == 0 {
		rst = append(rst, ValidateResult{line: 1, addMsg: "No labels found"})
	}

	return rst, nil
}
