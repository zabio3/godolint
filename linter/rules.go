package linter

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"path/filepath"
)

type Rule struct {
	Code     string
	Severity string
	Message  string
	CheckF   interface{} // func(node *parser.Node, file string) (rst []string, err error)
}

var RuleKeys = []string{
	"DL3000",
}

var Rules = map[string]*Rule{
	"DL3000": {
		Code:     "DL3000",
		Severity: "ErrorC",
		Message:  "Use absolute WORKDIR",
		CheckF:   DL3000Check,
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
