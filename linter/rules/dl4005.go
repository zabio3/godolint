package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// DL4005 Use SHELL to change the default shell
func dl4005Check(node *parser.Node, file string) (rst []string, err error) {

	for _, child := range node.Children {
		switch child.Value {
		case "run":
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
