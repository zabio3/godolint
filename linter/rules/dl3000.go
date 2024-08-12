package rules

import (
	"fmt"
	"path/filepath"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3000 is "Use absolute WORKDIR."
func validateDL3000(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == WORKDIR {
			absPath, err := filepath.Abs(child.Next.Value)
			if err != nil {
				return nil, fmt.Errorf("#%v DL3000: failed to convert relative path to absolute path (err: %s)", child.StartLine, err)
			}
			if absPath != child.Next.Value {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
