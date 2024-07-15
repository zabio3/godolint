package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4001 is dockerfile linter DL4001 rule.
// Either use Wget or Curl but not both.
func validateDL4001(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	var isCurl, isWget bool
	var lines []int
	for _, child := range node.Children {
		if child.Value == RUN {
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "curl":
					isCurl = true
					lines = append(lines, child.StartLine)
				case "wget":
					isWget = true
					lines = append(lines, child.StartLine)
				}
			}
		}
	}
	if isCurl && isWget {
		for _, line := range lines {
			rst = append(rst, ValidateResult{line: line})
		}
	}
	return rst, nil
}
