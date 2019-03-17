package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4004 Multiple `ENTRYPOINT` instructions found. If you list more than one `ENTRYPOINT` then only the last `ENTRYPOINT` will take effect
func validateDL4004(node *parser.Node, file string) (rst []string, err error) {
	isEntryPoint := false
	for _, child := range node.Children {
		switch child.Value {
		case "entrypoint":
			if !isEntryPoint {
				isEntryPoint = true
			} else {
				rst = append(rst, fmt.Sprintf("%s:%v DL4004 Multiple `ENTRYPOINT` instructions found. If you list more than one `ENTRYPOINT` then only the last `ENTRYPOINT` will take effect\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
