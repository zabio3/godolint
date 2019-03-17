package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
)

var regexDL3007 = regexp.MustCompile(`.*:latest`)

// validateDL3007 Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.
func validateDL3007(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "from" && regexDL3007.MatchString(child.Next.Value) {
			rst = append(rst, fmt.Sprintf("%s:%v DL3007 Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.\n", file, child.StartLine))
		}
	}
	return rst, nil
}
