package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3059 Multiple consecutive `RUN` instructions. Consider consolidation.
func validateDL3059(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	prevWasRun := false

	for _, child := range node.Children {
		if child.Value == RUN {
			if prevWasRun {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
			prevWasRun = true
		} else {
			prevWasRun = false
		}
	}
	return rst, nil
}
