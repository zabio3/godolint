package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"regexp"
	"strings"
)

var regexVersion3018 = regexp.MustCompile(`.+=.+`)

// validateDL3018 Do not use apk upgrade
func validateDL3018(node *parser.Node, file string) (rst []string, err error) {
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
				default:
					if isAdd && !regexVersion3018.MatchString(v) && length == len(rst) {
						rst = append(rst, fmt.Sprintf("%s:%v DL3018 Pin versions in apk add. Instead of `apk add <package>` use `apk add <package>=<version>`\n", file, child.StartLine))
						isApk, isAdd = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
