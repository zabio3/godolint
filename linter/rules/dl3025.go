package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// DL3025 Use arguments JSON notation for CMD and ENTRYPOINT arguments
func dl3025Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		isErr := false
		switch child.Value {
		case "entrypoint", "cmd":
			args := strings.Fields(child.Original)
			length := len(args) - 1
			for i, v := range strings.Fields(child.Original) {
				switch i {
				case 1:
					if v[:1] != "[" {
						isErr = true
					}
				case length:
					if v[len(v)-1:] != "]" {
						isErr = true
					}
				}
			}
			if isErr {
				rst = append(rst, fmt.Sprintf("%s:%v DL3025 Use arguments JSON notation for CMD and ENTRYPOINT arguments\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}
