package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3001 is "For some bash commands it makes no sense running them in a Docker container like ssh, vim, shutdown, service, ps, free, top, kill, mount, ifconfig."
func validateDL3001(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			for _, v := range strings.Fields(child.Next.Value) {
				for _, c := range []string{"ssh", "vim", "shutdown", "service", "ps", "free", "top", "kill", "mount"} {
					if v == c {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
					}
				}
			}
		}
	}
	return rst, nil
}
