package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// validateDL3024 FROM aliases (stage names) must be unique
func validateDL3024(node *parser.Node, file string) (rst []string, err error) {
	isAs := false
	var asBuildName []string
	for _, child := range node.Children {
		switch child.Value {
		case "from":
			for _, v := range strings.Fields(child.Original) {
				switch v {
				case "as":
					isAs = true
				default:
					if isAs {
						if isContains(asBuildName, v) {
							rst = append(rst, fmt.Sprintf("%s:%v DL3024 FROM aliases (stage names) must be unique\n", file, child.StartLine))
						} else {
							asBuildName = append(asBuildName, v)
						}
						isAs = false
					}
				}
			}
		}
	}
	return rst, nil
}
