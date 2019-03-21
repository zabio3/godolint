// Package linter provides dockerfile analyzer (Apply the rule to the parsed dockerfile).
package linter

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/zabio3/godolint/linter/rules"
)

// Analyzer Apply docker best practice rules to docker ast
func Analyzer(node *parser.Node, file string, ignoreRules []string) ([]string, error) {
	var rst []string
	for _, k := range GetMakeDifference(rules.RuleKeys, ignoreRules) {
		v, err := rules.Rules[k].ValidateFunc.(func(node *parser.Node, file string) (rst []string, err error))(node, file)
		if err != nil {
			return rst, err
		}
		rst = append(rst, v...)
	}
	return rst, nil
}

// GetMakeDifference is a function to create a difference set
func GetMakeDifference(xs, ys []string) []string {
	if len(xs) > len(ys) {
		return makeDifference(xs, ys)
	}
	return makeDifference(ys, xs)
}

// make set difference
func makeDifference(xs, ys []string) []string {
	var set []string
	for _, c := range xs {
		if !isContain(ys, c) {
			set = append(set, c)
		}
	}
	return set
}

// isContain is a function to check if s is in xs
func isContain(xs []string, s string) bool {
	for _, x := range xs {
		if s == x {
			return true
		}
	}
	return false
}
