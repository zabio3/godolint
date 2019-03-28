package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4001 Either use Wget or Curl but not both
func validateDL4001(node *parser.Node) (rst []ValidateResult, err error) {
	isCurl, isWget := false, false
	var numArr []int
	for _, child := range node.Children {
		switch child.Value {
		case RUN:
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "curl":
					isCurl = true
					numArr = append(numArr, child.StartLine)
				case "wget":
					isWget = true
					numArr = append(numArr, child.StartLine)
				}
			}
		}
		if isCurl && isWget {
			for _, num := range numArr {
				rst = append(rst, ValidateResult{line: num, addMsg: ""})
			}
		}
	}
	return rst, nil
}
