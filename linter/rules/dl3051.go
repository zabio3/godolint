package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3051 Label `<label>` is empty.
func validateDL3051(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)

			for key, val := range pairs {
				if strings.TrimSpace(val) == "" {
					rst = append(rst, ValidateResult{line: child.StartLine, addMsg: "Label has empty value: " + key})
					break
				}
			}
		}
	}
	return rst, nil
}
