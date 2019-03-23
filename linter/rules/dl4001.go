package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4001 Either use Wget or Curl but not both
func validateDL4001(node *parser.Node) (rst []ValidateResult, err error) {
	isCurl, isWget := false, false
	for _, child := range node.Children {
		switch child.Value {
		case RUN:
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "curl":
					isCurl = true
				case "wget":
					isWget = true
				}
			}
		}
		if isCurl && isWget {
			rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
		}
	}
	return rst, nil
}
