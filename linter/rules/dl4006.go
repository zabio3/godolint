package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4006 Set the `SHELL` option -o pipefail before `RUN` with a pipe in it
func validateDL4006(node *parser.Node) (rst []ValidateResult, err error) {
	var isShellPipeFail bool
	for _, child := range node.Children {
		switch child.Value {
		case SHELL:
			isShellPipeFail = true
		case RUN:
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "|":
					if !isShellPipeFail {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
				}
			}
		}
	}
	return rst, nil
}
