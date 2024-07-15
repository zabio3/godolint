package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4006 Set the `SHELL` option -o pipefail before `RUN` with a pipe in it
func validateDL4006(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	var isShellPipeFail bool
	for _, child := range node.Children {
		switch child.Value {
		case SHELL:
			isShellPipeFail = true
		case RUN:
			var isInQuote bool
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "|":
					if !isInQuote && !isShellPipeFail {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
				default:
					if strings.HasPrefix(v, "'") {
						isInQuote = true
					}
					if strings.HasSuffix(v, "'") {
						isInQuote = false
					}
				}
			}
		}
	}
	return rst, nil
}
