package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
)

var verPattern3007 = regexp.MustCompile(`.*:latest`)

// dl3007 Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.
func dl3007Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "from" {
			if verPattern3007.MatchString(child.Next.Value) {
				rst = append(rst, fmt.Sprintf("%s:%v DL3007 Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
