package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3047 `wget` without flag `--progress` will result in excessively bloated build logs when downloading larger files.
func validateDL3047(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			fields := strings.Fields(child.Next.Value)
			for i := 0; i < len(fields); i++ {
				if fields[i] == "wget" {
					hasFlag := false
					for j := i + 1; j < len(fields) && fields[j] != "&&"; j++ {
						if strings.HasPrefix(fields[j], "--progress") {
							hasFlag = true
							break
						}
					}
					if !hasFlag {
						rst = append(rst, ValidateResult{line: child.StartLine})
						break
					}
				}
			}
		}
	}
	return rst, nil
}
