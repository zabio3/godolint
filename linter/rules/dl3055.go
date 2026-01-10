package rules

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Git hash is 40 hexadecimal characters
var gitHashPattern3055 = regexp.MustCompile(`^[a-f0-9]{40}$`)

// Labels that should contain git hashes
var gitHashLabels3055 = []string{
	"org.opencontainers.image.revision",
}

// validateDL3055 Label `<label>` is not a valid git hash.
func validateDL3055(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)

			for key, val := range pairs {
				isGitHashLabel := false
				for _, hashLabel := range gitHashLabels3055 {
					if key == hashLabel {
						isGitHashLabel = true
						break
					}
				}

				if isGitHashLabel {
					if !gitHashPattern3055.MatchString(val) {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: "Label value is not a valid git hash: " + key})
						break
					}
				}
			}
		}
	}
	return rst, nil
}
