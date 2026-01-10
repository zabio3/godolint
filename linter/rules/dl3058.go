package rules

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Simplified RFC5322 email pattern
var emailPattern3058 = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Labels that should contain email addresses
var emailLabels3058 = []string{
	"org.opencontainers.image.authors",
	"maintainer",
}

// validateDL3058 Label `<label>` is not a valid RFC5322 email format.
func validateDL3058(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)

			for key, val := range pairs {
				isEmailLabel := false
				for _, emailLabel := range emailLabels3058 {
					if key == emailLabel {
						isEmailLabel = true
						break
					}
				}

				if isEmailLabel {
					if !emailPattern3058.MatchString(val) {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: "Label value is not a valid email: " + key})
						break
					}
				}
			}
		}
	}
	return rst, nil
}
