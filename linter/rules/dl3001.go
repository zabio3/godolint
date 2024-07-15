package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3001 is "For some bash commands it makes no sense running them in a Docker container
// like free, ifconfig, kill, mount, ps, service, shutdown, ssh, top, vim"
func validateDL3001(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			for _, v := range strings.Fields(child.Next.Value) {
				if existsTools(v) {
					rst = append(rst, ValidateResult{line: child.StartLine})
				}
			}
		}
	}
	return rst, nil
}

func existsTools(s string) bool {
	for _, c := range []string{"free", "ifconfig", "kill", "mount", "ps", "service", "shutdown", "ssh", "top", "vim"} {
		if s == c {
			return true
		}
	}
	return false
}
