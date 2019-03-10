package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// dl3004Check is "Do not use sudo as it leads to unpredictable behavior. Use a tool like gosu to enforce root."
func dl3004Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			for _, v := range strings.Fields(child.Next.Value) {
				if v == "sudo" {
					rst = append(rst, fmt.Sprintf("%s:%v DL3004 Do not use sudo as it leads to unpredictable behavior. Use a tool like gosu to enforce root.\n", file, child.StartLine))
				}
			}
		}
	}
	return rst, nil
}
