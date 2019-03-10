package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
)

// dl3006 Always tag the version of an image explicitly"
func dl3006Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "from" && !regexp.MustCompile(`.+[:].+`).MatchString(child.Next.Value) {
			rst = append(rst, fmt.Sprintf("%s:%v DL3006 Always tag the version of an image explicitly\n", file, child.StartLine))
		}
	}
	return rst, nil
}
