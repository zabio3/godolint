package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// validateDL3003 is "Use WORKDIR to switch to a directory"
func validateDL3003(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			for _, v := range strings.Fields(child.Next.Value) {
				if v == "cd" {
					rst = append(rst, fmt.Sprintf("%s:%v DL3003 Use WORKDIR to switch to a directory\n", file, child.StartLine))
				}
			}
		}
	}
	return rst, nil
}
