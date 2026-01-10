package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Common SPDX license identifiers
var spdxLicenses3054 = map[string]bool{
	"MIT":          true,
	"Apache-2.0":   true,
	"GPL-3.0":      true,
	"GPL-2.0":      true,
	"BSD-3-Clause": true,
	"BSD-2-Clause": true,
	"ISC":          true,
	"LGPL-3.0":     true,
	"LGPL-2.1":     true,
	"MPL-2.0":      true,
	"AGPL-3.0":     true,
	"Unlicense":    true,
	"CC0-1.0":      true,
	"EPL-2.0":      true,
	"EPL-1.0":      true,
}

// Labels that should contain SPDX license identifiers
var licenseLabels3054 = []string{
	"org.opencontainers.image.licenses",
}

// validateDL3054 Label `<label>` is not a valid SPDX license identifier.
func validateDL3054(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)

			for key, val := range pairs {
				isLicenseLabel := false
				for _, licLabel := range licenseLabels3054 {
					if key == licLabel {
						isLicenseLabel = true
						break
					}
				}

				if isLicenseLabel {
					if !spdxLicenses3054[val] {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: "Label value is not a valid SPDX license: " + key})
						break
					}
				}
			}
		}
	}
	return rst, nil
}
