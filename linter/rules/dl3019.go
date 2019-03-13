package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// dl3019 Use the `--no-cache` switch to avoid the need to use `--update` and remove `/var/cache/apk/*` when done installing packages
func dl3019Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			isApk, isAdd, length := false, false, len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apk":
					isApk = true
				case "add":
					if isApk {
						isAdd = true
					}
				case "&&":
					isApk, isAdd = false, false
					continue
				default:
					if isAdd && v != "--update" && length == len(rst) {
						rst = append(rst, fmt.Sprintf("%s:%v DL3019 Use the `--no-cache` switch to avoid the need to use `--update` and remove `/var/cache/apk/*` when done installing packages\n", file, child.StartLine))
						isApk, isAdd = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
