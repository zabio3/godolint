package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// DL3023 COPY --from should reference a previously defined FROM alias
func dl3023Check(node *parser.Node, file string) (rst []string, err error) {
	fromImage := ""
	isAs, isAsBuild, isMultiFrom, isFromBuild := false, false, false, false

	for _, child := range node.Children {
		switch child.Value {
		case "from":
			for _, v := range strings.Fields(child.Original) {
				switch v {
				case "as":
					isAs = true
				case "build":
					if isAs {
						isAsBuild = true
					}
				default:
					if fromImage == "" && v != "FROM" && v != "from" {
						fromImage = v
					} else if fromImage == v && isAsBuild {
						isMultiFrom = true
					}
				}
			}
		case "copy":
			for _, v := range strings.Fields(child.Original) {
				if v == "--from=build" {
					isFromBuild = true
				}
			}
			if isAsBuild && !isMultiFrom && isFromBuild {
				rst = append(rst, fmt.Sprintf("%s:%v DL3023 COPY --from should reference a previously defined FROM alias\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
