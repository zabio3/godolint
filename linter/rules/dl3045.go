package rules

import (
	"path/filepath"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3045 `COPY` to a relative destination without `WORKDIR` set.
func validateDL3045(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	hasWorkdir := false

	for _, child := range node.Children {
		switch child.Value {
		case WORKDIR:
			hasWorkdir = true
		case FROM:
			// Reset WORKDIR state for each stage
			hasWorkdir = false
		case COPY:
			if !hasWorkdir {
				// Get the destination (last argument of COPY)
				dest := getLastArg(child.Next)
				if dest != "" && !filepath.IsAbs(dest) && dest != "." && !strings.HasPrefix(dest, "./") {
					rst = append(rst, ValidateResult{line: child.StartLine})
				}
			}
		}
	}
	return rst, nil
}

// getLastArg gets the last argument from a linked list of nodes
func getLastArg(node *parser.Node) string {
	if node == nil {
		return ""
	}

	var last string
	for n := node; n != nil; n = n.Next {
		last = n.Value
	}
	return last
}
