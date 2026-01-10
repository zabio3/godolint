package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3043 `ONBUILD`, `FROM` or `MAINTAINER` triggered from within `ONBUILD` instruction is not allowed.
func validateDL3043(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == ONBUILD {
			// The nested instruction is in child.Next.Children
			if child.Next != nil && len(child.Next.Children) > 0 {
				nestedInstruction := child.Next.Children[0].Value
				if nestedInstruction == FROM || nestedInstruction == MAINTAINER || nestedInstruction == ONBUILD {
					rst = append(rst, ValidateResult{line: child.StartLine})
				}
			}
		}
	}
	return rst, nil
}
