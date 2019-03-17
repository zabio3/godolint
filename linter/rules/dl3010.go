package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
	"strings"
)

var isCompressionExt = regexp.MustCompile(`(?:\.tar\.gz|tar\.bz|\.tar.xz|\.tgz|\.tbz)$`)

// validateDL3010 Use ADD for extracting archives into an image.
func validateDL3010(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "copy" {
			args := strings.Fields(child.Next.Value)
			if len(args) >= 1 && isCompressionExt.MatchString(args[0]) {
				rst = append(rst, fmt.Sprintf("%s:%v DL3010 Use ADD for extracting archives into an image.\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
