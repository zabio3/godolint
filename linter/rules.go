package linter

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"path/filepath"
	"strings"
)

type Rule struct {
	Code     string
	Severity string
	CheckF   interface{} // func(node *parser.Node, file string) (rst []string, err error)
}

var RuleKeys = []string{
	"DL3000",
	"DL3001",
}

var Rules = map[string]*Rule{
	"DL3000": {
		Code:     "DL3000",
		Severity: "ErrorC",
		CheckF:   DL3000Check,
	},
	"DL3001": {
		Code:     "DL3001",
		Severity: "InfoC",
		CheckF:   DL3001Check,
	},
}

func DL3000Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "workdir" {
			absPath, err := filepath.Abs(child.Next.Value)
			if err != nil {
				return nil, err
			}
			if absPath != child.Next.Value {
				rst = append(rst, fmt.Sprintf("%s:%v DL3000 Use absolute WORKDIR\n", file, child.StartLine))
			}
		}
	}
	return rst, nil
}

func DL3001Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			for _, v := range strings.Fields(child.Next.Value) {
				for _, c := range []string{"ssh", "vim", "shutdown", "service", "ps", "free", "top", "kill", "mount"} {
					if v == c {
						rst = append(rst, fmt.Sprintf("%s:%v DL3001 For some bash commands it makes no sense running them in a Docker container like `ssh`, `vim`, `shutdown`, `service`, `ps`, `free`, `top`, `kill`, `mount`, `ifconfig`\n", file, child.StartLine))
					}
				}
			}
		}
	}
	return rst, nil
}
