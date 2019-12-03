package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4004 Multiple `ENTRYPOINT` instructions found. If you list more than one `ENTRYPOINT` then only the last `ENTRYPOINT` will take effect
func validateDL4004(node *parser.Node) (rst []ValidateResult, err error) {
	isEntryPoint := false
	for _, child := range node.Children {
		if child.Value == ENTRYPOINT {
			if !isEntryPoint {
				isEntryPoint = true
			} else {
				rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
			}
		}
	}
	return rst, nil
}
