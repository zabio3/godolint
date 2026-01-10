package rules

import (
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Label key format: [a-z0-9]+([._-][a-z0-9]+)*
var labelKeyPattern = regexp.MustCompile(`^[a-z0-9]+([._-][a-z0-9]+)*(/[a-z0-9]+([._-][a-z0-9]+)*)*$`)

// validateDL3048 Invalid label key.
func validateDL3048(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == LABEL {
			pairs := parseLabelPairsFromNode(child.Next)

			for key := range pairs {
				if !labelKeyPattern.MatchString(key) {
					rst = append(rst, ValidateResult{line: child.StartLine, addMsg: "Invalid label key: " + key})
					break
				}
			}
		}
	}
	return rst, nil
}

// parseLabelPairsFromNode extracts key-value pairs from LABEL instruction node
func parseLabelPairsFromNode(n *parser.Node) map[string]string {
	pairs := make(map[string]string)

	// Walk the linked list: key -> value -> = -> key -> value -> ...
	for n != nil {
		key := n.Value
		if n.Next != nil && n.Next.Value != "=" {
			value := n.Next.Value
			// Strip quotes from value
			value = strings.Trim(value, `"'`)
			pairs[key] = value
			// Skip to next pair (past the "=" node if present)
			n = n.Next.Next
			if n != nil && n.Value == "=" {
				n = n.Next
			}
		} else {
			n = n.Next
		}
	}

	return pairs
}
