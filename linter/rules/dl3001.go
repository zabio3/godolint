package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// validateDL3001 is "For some bash commands it makes no sense running them in a Docker container like ssh, vim, shutdown, service, ps, free, top, kill, mount, ifconfig."
func validateDL3001(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			for _, v := range strings.Fields(child.Next.Value) {
				for _, c := range []string{"ssh", "vim", "shutdown", "service", "ps", "free", "top", "kill", "mount"} {
					if v == c {
						rst = append(rst, fmt.Sprintf("%s:%v DL3001 For some bash commands it makes no sense running them in a Docker container like `ssh`, `vim`, `shutdown`, `service`, `ps`, `free`, `top`, `kill`, `mount`, `ifconfig`\n", file, child.StartLine))
					}
				}
			}
		}
	}
	return rst, nil
}
