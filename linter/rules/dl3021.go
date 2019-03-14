package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// DL3021 `COPY` with more than 2 arguments requires the last argument to end with `/`
func dl3021Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "copy" {
			if isDL3021Error(child) {
				rst = append(rst, fmt.Sprintf("%s:%v DL3021 `COPY` with more than 2 arguments requires the last argument to end with `/`\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}

func isDL3021Error(node *parser.Node) bool {
	c := 0
	var fn func(nd *parser.Node, str string) bool
	fn = func(nd *parser.Node, str string) bool {
		switch nd {
		case nil:
			if c > 3 && str[len(str)-1:] != "/" {
				return true
			}
			return false
		default:
			c += 1
			return fn(nd.Next, nd.Value)
		}
	}
	return fn(node, "")
}
