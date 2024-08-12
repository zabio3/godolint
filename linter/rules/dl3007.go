package rules

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexDL3007 = regexp.MustCompile(`.*:latest`)

// validateDL3007 Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.
func validateDL3007(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == FROM && regexDL3007.MatchString(child.Next.Value) {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}
