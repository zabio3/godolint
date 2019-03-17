package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
)

var regexDL3006 = regexp.MustCompile(`.+[:].+`)

// validateDL3006 Always tag the version of an image explicitly"
func validateDL3006(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "from" && !regexDL3006.MatchString(child.Next.Value) {
			rst = append(rst, fmt.Sprintf("%s:%v DL3006 Always tag the version of an image explicitly\n", file, child.StartLine))
		}
	}
	return rst, nil
}
