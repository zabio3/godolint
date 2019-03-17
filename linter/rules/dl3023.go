package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
	"strings"
)

var regexVersion3023 = regexp.MustCompile(`--from=.+`)

// validateDL3023 COPY --from should reference a previously defined FROM alias
func validateDL3023(node *parser.Node, file string) (rst []string, err error) {
	isAs := false
	asFromName := ""
	for _, child := range node.Children {
		switch child.Value {
		case "from":
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
		case "copy":
			for _, v := range strings.Fields(child.Original) {
				if regexVersion3023.MatchString(v) && v == fmt.Sprintf("--from=%s", asFromName) {
					rst = append(rst, fmt.Sprintf("%s:%v DL3023 COPY --from should reference a previously defined FROM alias\n", file, child.StartLine))
				}
			}
		}
	}
	return rst, nil
}
