package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
	"strings"
)

// dl3010 Use ADD for extracting archives into an image.
var isTar = regexp.MustCompile(`(?:\.tar\.gz|tar\.bz|\.tar.xz|\.tgz|\.tbz)$`)

func dl3010Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "copy" {
			args := strings.Fields(child.Next.Value)
			if len(args) >= 1 && isTar.MatchString(args[0]) {
				rst = append(rst, fmt.Sprintf("%s:%v DL3010 Use ADD for extracting archives into an image.\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
