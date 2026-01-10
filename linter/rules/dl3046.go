package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3046 `useradd` without flag `-l` and target UID greater than or equal to 65534 can lead to excessively large Image.
func validateDL3046(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			fields := strings.Fields(child.Next.Value)
			for i := 0; i < len(fields); i++ {
				if fields[i] == "useradd" {
					hasFlag := false
					for j := i + 1; j < len(fields) && fields[j] != "&&"; j++ {
						v := fields[j]
						if v == "-l" || v == "--no-log-init" {
							hasFlag = true
							break
						}
						// Check for combined short flags like -ml
						if strings.HasPrefix(v, "-") && !strings.HasPrefix(v, "--") && strings.Contains(v, "l") {
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
