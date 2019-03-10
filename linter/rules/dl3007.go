package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
)

// dl3007 Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.
func dl3007Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "from" {
			matched, _ := regexp.MatchString(`.*:latest`, child.Next.Value)
			//if err != nil {
			//	return rst, err
			//}
			if matched {
				rst = append(rst, fmt.Sprintf("%s:%v DL3007 Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
