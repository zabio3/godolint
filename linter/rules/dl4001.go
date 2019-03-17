package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// DL4001 Either use Wget or Curl but not both
func dl4001Check(node *parser.Node, file string) (rst []string, err error) {
	isCurl, isWget := false, false
	for _, child := range node.Children {
		fmt.Println(child)
		switch child.Value {
		case "run":
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
			rst = append(rst, fmt.Sprintf("%s:%v DL4001 Either use Wget or Curl but not both\n", file, child.StartLine))
		}
	}
	return rst, nil
}
