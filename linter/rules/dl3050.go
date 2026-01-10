package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Standard OCI labels that are generally allowed
var allowedLabelPrefixes = []string{
	"org.opencontainers.image",
	"com.example",
	"version",
	"maintainer",
	"description",
}

// validateDL3050 Superfluous label(s) present.
func validateDL3050(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)

			for key := range pairs {
				allowed := false
				for _, prefix := range allowedLabelPrefixes {
					if strings.HasPrefix(key, prefix) || key == prefix {
						allowed = true
						break
					}
				}
				if !allowed {
					rst = append(rst, ValidateResult{line: child.StartLine, addMsg: "Label not in allowlist: " + key})
					break
				}
			}
		}
	}
	return rst, nil
}
