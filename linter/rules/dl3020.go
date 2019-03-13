package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// DL3020 Use COPY instead of ADD for files and folders
func dl3020Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "add" {
			rst = append(rst, fmt.Sprintf("%s:%v DL3020 Use COPY instead of ADD for files and folders\n", file, child.StartLine))
		}
	}
	return rst, nil
}
