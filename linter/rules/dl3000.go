package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"path/filepath"
)

// dl3000Check is "Use absolute WORKDIR."
func dl3000Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "workdir" {
			absPath, _ := filepath.Abs(child.Next.Value)
			//if err != nil {
			//	return nil, err
			//}
			if absPath != child.Next.Value {
				rst = append(rst, fmt.Sprintf("%s:%v DL3000 Use absolute WORKDIR\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
