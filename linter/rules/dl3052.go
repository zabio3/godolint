package rules

import (
	"net/url"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Labels that should contain URLs
var urlLabels3052 = []string{
	"org.opencontainers.image.url",
	"org.opencontainers.image.documentation",
	"org.opencontainers.image.source",
}

// validateDL3052 Label `<label>` is not a valid URL.
func validateDL3052(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)

			for key, val := range pairs {
				isURLLabel := false
				for _, urlLabel := range urlLabels3052 {
					if key == urlLabel {
						isURLLabel = true
						break
					}
				}

				if isURLLabel {
					if _, err := url.ParseRequestURI(val); err != nil {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: "Label value is not a valid URL: " + key})
						break
					}
				}
			}
		}
	}
	return rst, nil
}
