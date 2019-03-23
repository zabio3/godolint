package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3025 Use arguments JSON notation for CMD and ENTRYPOINT arguments
func validateDL3025(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		isErr := false
		switch child.Value {
		case ENTRYPOINT, CMD:
			args := strings.Fields(child.Original)
			length := len(args) - 1
			for i, v := range strings.Fields(child.Original) {
				switch i {
				case 1:
					if v[:1] != "[" {
						isErr = true
					}
				case length:
					if v[len(v)-1:] != "]" {
						isErr = true
					}
				}
			}
			if isErr {
				rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
			}
		}
	}
	return rst, nil
}
