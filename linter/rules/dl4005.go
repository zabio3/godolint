package rules

import (
	"fmt"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4005 Use SHELL to change the default shell
func validateDL4005(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		switch child.Value {
		case RUN:
			isLn := false
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "ln":
					isLn = true
				case "/bin/sh":
					if isLn {
						rst = append(rst, fmt.Sprintf("%s:%v DL4005 Use SHELL to change the default shell\n", file, child.StartLine))
					}
				}
			}
		}
	}
	return rst, nil
}
