package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3021 `COPY` with more than 2 arguments requires the last argument to end with `/`
func validateDL3021(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == COPY {
			if isDL3021Error(child) {
				rst = append(rst, ValidateResult{line: child.StartLine})
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
			c++
			return fn(nd.Next, nd.Value)
		}
	}
	return fn(node, "")
}
