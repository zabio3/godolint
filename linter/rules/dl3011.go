package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strconv"
)

// validateDL3011 Valid UNIX ports range from 0 to 65535
func validateDL3011(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "expose" {
			port := child.Next
			if port != nil {
				portNum, err := strconv.Atoi(port.Value)
				if err != nil {
					return nil, fmt.Errorf("%s:%v DL3011 not numeric is the value set for the port: %s", file, child.StartLine, port.Value)
				}
				if portNum < 0 || portNum > 65535 {
					rst = append(rst, fmt.Sprintf("%s:%v DL3011 Valid UNIX ports range from 0 to 65535\n", file, child.StartLine))
				}
			}
		}
	}
	return rst, nil
}
