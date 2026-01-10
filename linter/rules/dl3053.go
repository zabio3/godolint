package rules

import (
	"time"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Labels that should contain RFC3339 datetime
var datetimeLabels3053 = []string{
	"org.opencontainers.image.created",
}

// validateDL3053 Label `<label>` is not a valid RFC3339 format datetime.
func validateDL3053(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)

			for key, val := range pairs {
				isDatetimeLabel := false
				for _, dtLabel := range datetimeLabels3053 {
					if key == dtLabel {
						isDatetimeLabel = true
						break
					}
				}

				if isDatetimeLabel {
					if _, err := time.Parse(time.RFC3339, val); err != nil {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: "Label value is not RFC3339 format: " + key})
						break
					}
				}
			}
		}
	}
	return rst, nil
}
