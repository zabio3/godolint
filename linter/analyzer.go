// Package linter provides dockerfile analyzer (Apply the rule to the parsed dockerfile).
package linter

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/zabio3/godolint/linter/rules"
)

// Analyzer implements Analyzer.
type Analyzer struct {
	rules []*rules.Rule
}

// NewAnalyzer generate a NewAnalyzer with rules to apply
func NewAnalyzer(ignoreRules []string) Analyzer {
	return newAnalyzer(ignoreRules)
}

func newAnalyzer(ignoreRules []string) Analyzer {
	var filteredRules []*rules.Rule
	for _, k := range getMakeDifference(rules.RuleKeys, ignoreRules) {
		if rule, ok := rules.Rules[k]; ok {
			filteredRules = append(filteredRules, rule)
		}
	}
	return Analyzer{rules: filteredRules}
}

// Run apply docker best practice rules to docker ast
func (a Analyzer) Run(node *parser.Node) ([]string, error) {
	var rst []string
	rstChan := make(chan []string, len(a.rules))
	errChan := make(chan error, len(a.rules))

	for _, rule := range a.rules {
		go func(r *rules.Rule) {
			vrst, err := r.ValidateFunc(node)
			if err != nil {
				errChan <- err
			} else {
				rstChan <- rules.CreateMessage(rule, vrst)
			}
		}(rule)
		select {
		case value := <-rstChan:
			rst = append(rst, value...)
		case err := <-errChan:
			return nil, err
		}
	}

	return rst, nil
}

// getMakeDifference is a function to create a difference set
func getMakeDifference(xs, ys []string) []string {
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
