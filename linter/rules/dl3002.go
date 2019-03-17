package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3002 is "Last user should not be root."
func validateDL3002(node *parser.Node, file string) (rst []string, err error) {
	isLastRootUser := false
	var lastRootUserPos int
	for _, child := range node.Children {
		if child.Value == "user" {
			if child.Next.Value == "root" || child.Next.Value == "0" {
				isLastRootUser = true
				lastRootUserPos = child.StartLine
			} else {
				isLastRootUser = false
				lastRootUserPos = 0
			}
		}
	}
	if isLastRootUser {
		rst = append(rst, fmt.Sprintf("%s:%v DL3002 Last USER should not be root\n", file, lastRootUserPos))
	}

	return rst, nil
}
