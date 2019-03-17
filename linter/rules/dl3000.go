package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"path/filepath"
)

// validateDL3000 is "Use absolute WORKDIR."
func validateDL3000(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "workdir" {
			absPath, err := filepath.Abs(child.Next.Value)
			if err != nil {
				return nil, err
			}
			if absPath != child.Next.Value {
				rst = append(rst, fmt.Sprintf("%s:%v DL3000 Use absolute WORKDIR\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
