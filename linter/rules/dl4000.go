package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4000 MAINTAINER is deprecated
func validateDL4000(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		switch child.Value {
		case MAINTAINER:
			rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
		}
	}
	return rst, nil
}
