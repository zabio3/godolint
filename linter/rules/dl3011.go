package rules

import (
	"fmt"
	"strconv"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3011 Valid UNIX ports range from 0 to 65535
func validateDL3011(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == EXPOSE {
			port := child.Next
			if port != nil {
				portNum, err := strconv.Atoi(port.Value)
				if err != nil {
					return nil, fmt.Errorf("#%v DL3011 not numeric is the value set for the port: %s", child.StartLine, port.Value)
				}
				if portNum < 0 || portNum > 65535 {
					rst = append(rst, ValidateResult{line: child.StartLine})
				}
			}
		}
	}
	return rst, nil
}
