package rules

import (
	"path/filepath"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3000 is "Use absolute WORKDIR."
func validateDL3000(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == WORKDIR {
			absPath, err := filepath.Abs(child.Next.Value)
			if err != nil {
				return nil, err
			}
			if absPath != child.Next.Value {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
