package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var isCompressionExt = regexp.MustCompile(`(?:\.tar\.gz|tar\.bz|\.tar.xz|\.tgz|\.tbz)$`)

// validateDL3010 Use ADD for extracting archives into an image.
func validateDL3010(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == COPY {
			args := strings.Fields(child.Next.Value)
			if len(args) >= 1 && isCompressionExt.MatchString(args[0]) {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
