package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3012 Multiple `HEALTHCHECK` instructions found. Only the last `HEALTHCHECK` will take effect.
func validateDL3012(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	if node == nil {
		return rst, nil
	}
	var healthcheckCount int
	for _, child := range node.Children {
		if child.Value == HEALTHCHECK {
			healthcheckCount++
			if healthcheckCount > 1 {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
