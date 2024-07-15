package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexVersion3023 = regexp.MustCompile(`--from=.+`)

// validateDL3023 COPY --from should reference a previously defined FROM alias
func validateDL3023(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	var isAs bool
	asFromName := ""
	for _, child := range node.Children {
		switch child.Value {
		case FROM:
			for _, v := range strings.Fields(child.Original) {
				switch v {
				case "as":
					isAs = true
				default:
					if isAs {
						asFromName = v
						isAs = false
					} else {
						asFromName = ""
					}
				}
			}
		case COPY:
			for _, v := range strings.Fields(child.Original) {
				if regexVersion3023.MatchString(v) && v == fmt.Sprintf("--from=%s", asFromName) {
					rst = append(rst, ValidateResult{line: child.StartLine})
				}
			}
		}
	}
	return rst, nil
}
