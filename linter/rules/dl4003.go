package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// DL4003 Either use Wget or Curl but not both
func dl4003Check(node *parser.Node, file string) (rst []string, err error) {
	isCmd := false
	for _, child := range node.Children {
		switch child.Value {
		case "cmd":
			if !isCmd {
				isCmd = true
			} else {
				rst = append(rst, fmt.Sprintf("%s:%v DL4003 Multiple `CMD` instructions found. If you list more than one `CMD` then only the last `CMD` will take effect\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
