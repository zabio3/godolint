package linter

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"path/filepath"
	"strings"
)

// Rule is filtered rule (with ignore rule applied)
// CheckF func(node *parser.Node, file string) (rst []string, err error)
type Rule struct {
	Code     string
	Severity string
	CheckF   interface{}
}

// RuleKeys is (Docker best practice rule key)
var RuleKeys = []string{
	"DL3000",
	"DL3001",
	"DL3002",
	"DL3003",
}

// Rules (Docker best practice rule key)
var Rules = map[string]*Rule{
	"DL3000": {
		Code:     "DL3000",
		Severity: "ErrorC",
		CheckF:   dL3000Check,
	},
	"DL3001": {
		Code:     "DL3001",
		Severity: "InfoC",
		CheckF:   dL3001Check,
	},
	"DL3002": {
		Code:     "DL3002",
		Severity: "WarningC",
		CheckF:   dL3002Check,
	},
	"DL3003": {
		Code:     "DL3003",
		Severity: "WarningC",
		CheckF:   dL3003Check,
	},
}

// dL3000Check is "Use absolute WORKDIR."
func dL3000Check(node *parser.Node, file string) (rst []string, err error) {
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

// dL3001Check is "For some bash commands it makes no sense running them in a Docker container like ssh, vim, shutdown, service, ps, free, top, kill, mount, ifconfig."
func dL3001Check(node *parser.Node, file string) (rst []string, err error) {
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

// dL3002Check is "Last user should not be root."
func dL3002Check(node *parser.Node, file string) (rst []string, err error) {
	var isLastRootUser = false
	var lastRootUserPos int
	for _, child := range node.Children {
		if child.Value == "user" {
			if child.Next.Value == "root" || child.Next.Value == "0" {
				isLastRootUser = true
				lastRootUserPos = child.StartLine
			} else {
				isLastRootUser = false
				lastRootUserPos = 0
			}
		}
	}
	if isLastRootUser {
		rst = append(rst, fmt.Sprintf("%s:%v DL3002 Last USER should not be root\n", file, lastRootUserPos))
		return rst, nil
	}

	return rst, nil
}

// dL3003Check is "Use WORKDIR to switch to a directory"
func dL3003Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		if child.Value == "run" {
			for _, v := range strings.Fields(child.Next.Value) {
				if v == "cd" {
					rst = append(rst, fmt.Sprintf("%s:%v DL3003 Use WORKDIR to switch to a directory\n", file, child.StartLine))
				}
			}
		}
	}
	return rst, nil
}
