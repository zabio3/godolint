package rules

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Semantic version pattern
var semverPattern3056 = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// Labels that should contain semantic versions
var versionLabels3056 = []string{
	"org.opencontainers.image.version",
}

// validateDL3056 Label `<label>` does not conform to semantic versioning.
func validateDL3056(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)

			for key, val := range pairs {
				isVersionLabel := false
				for _, verLabel := range versionLabels3056 {
					if key == verLabel {
						isVersionLabel = true
						break
					}
				}

				if isVersionLabel {
					if !semverPattern3056.MatchString(val) {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: "Label value is not a valid semantic version: " + key})
						break
					}
				}
			}
		}
	}
	return rst, nil
}
