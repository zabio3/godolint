package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4003 Either use Wget or Curl but not both
func validateDL4003(node *parser.Node) (rst []ValidateResult, err error) {
	isCmd := false
	for _, child := range node.Children {
		switch child.Value {
		case CMD:
			if !isCmd {
				isCmd = true
			} else {
				rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
			}
		}
	}
	return rst, nil
}
