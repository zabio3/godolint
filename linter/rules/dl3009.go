package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// dl3009 Delete the apt-get lists after installing something.
func dl3009Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			isInstall, isCleanAptDone := false, false
			args := strings.Fields(child.Next.Value)
			for i, v := range args {
				switch v {
				case "apt-get":
					if len(args) > i+1 {
						switch args[i+1] {
						case "update", "install":
							isInstall = true
						default:
							if isInstall {
								isCleanAptDone = hasCleanApt(args[i+1:])
							}
						}
					}
				}
			}
			if isInstall && !isCleanAptDone {
				rst = append(rst, fmt.Sprintf("%s:%v DL3009 Delete the apt-get lists after installing something\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}

func hasCleanApt(args []string) bool {
	hasClean, hasRemove := false, false
	size := len(args)
	for i, v := range args {
		switch v {
		case "&&":
			continue
		case "apt-get":
			if size > i+1 && args[i+1] == "clean" {
				hasClean = true
			}
		case "rm":
			if size > i+1 {
				for _, w := range args[i+1:] {
					switch w {
					case "&&":
						continue
					case "/var/lib/apt/lists/*":
						hasRemove = true
					}
				}
			}
		}
	}
	return hasClean && hasRemove
}
