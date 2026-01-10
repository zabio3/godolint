package rules

import (
	"slices"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// unsupportedTools is a list of commands that make no sense in a Docker container.
var unsupportedTools = []string{"free", "ifconfig", "kill", "mount", "ps", "service", "shutdown", "ssh", "top", "vim"}

// validateDL3001 validates that certain bash commands are not used in Docker containers.
// Commands like free, ifconfig, kill, mount, ps, service, shutdown, ssh, top, vim
// make no sense running in a Docker container.
func validateDL3001(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			for _, v := range strings.Fields(child.Next.Value) {
				if slices.Contains(unsupportedTools, v) {
					rst = append(rst, ValidateResult{line: child.StartLine})
				}
			}
		}
	}
	return rst, nil
}
