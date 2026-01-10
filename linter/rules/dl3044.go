package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3044 Do not refer to an environment variable within the same ENV statement where it is defined.
func validateDL3044(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == ENV {
			// Walk the linked list to collect key-value pairs
			// ENV format: key -> value -> = -> key -> value -> = ...
			var pairs []envPair
			n := child.Next
			for n != nil {
				key := n.Value
				if n.Next != nil && n.Next.Value != "=" {
					value := n.Next.Value
					pairs = append(pairs, envPair{key: key, value: value})
					// Skip to next pair (past the "=" node if present)
					n = n.Next.Next
					if n != nil && n.Value == "=" {
						n = n.Next
					}
				} else {
					// Simple KEY=VALUE in the key itself
					if strings.Contains(key, "=") {
						kv := strings.SplitN(key, "=", 2)
						pairs = append(pairs, envPair{key: kv[0], value: kv[1]})
					}
					n = n.Next
				}
			}

			// Check if any defined variable references itself in its value
			for _, pair := range pairs {
				// Check for $VAR or ${VAR} references using string matching
				if containsSelfReference(pair.key, pair.value) {
					rst = append(rst, ValidateResult{line: child.StartLine})
					break
				}
			}
		}
	}
	return rst, nil
}

// containsSelfReference checks if value contains a reference to the variable key
// by looking for $KEY or ${KEY} patterns
func containsSelfReference(key, value string) bool {
	// Check for ${KEY} pattern
	if strings.Contains(value, "${"+key+"}") {
		return true
	}

	// Check for $KEY pattern (must be followed by non-identifier char or end of string)
	dollarKey := "$" + key
	searchStart := 0
	for {
		idx := strings.Index(value[searchStart:], dollarKey)
		if idx == -1 {
			return false
		}
		idx += searchStart // Convert to absolute index
		endPos := idx + len(dollarKey)

		// $KEY at end of string is a match
		if endPos >= len(value) {
			return true
		}

		// Check if next char is not a valid identifier character
		if !isIdentifierChar(value[endPos]) {
			return true
		}

		// Continue searching from after this occurrence
		searchStart = endPos
	}
}

// isIdentifierChar returns true if c is a valid identifier character (a-zA-Z0-9_)
func isIdentifierChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_'
}

type envPair struct {
	key   string
	value string
}
