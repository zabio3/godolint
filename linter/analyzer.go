// Package linter provides dockerfile analyzer (Apply the rule to the parsed dockerfile).
package linter

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"

	"github.com/zabio3/godolint/linter/rules"
)

// Analyzer implements Analyzer.
type Analyzer struct {
	rules             []*rules.Rule
	trustedRegistries []string
}

// NewAnalyzer generate a NewAnalyzer with rules to apply.
func NewAnalyzer(ignoreRules []string, trustedRegistries []string) Analyzer {
	return newAnalyzer(ignoreRules, trustedRegistries)
}

func newAnalyzer(ignoreRules []string, trustedRegistries []string) Analyzer {
	var filteredRules []*rules.Rule
	for _, k := range getMakeDiff(rules.RuleKeys, ignoreRules) {
		if rule, ok := rules.Rules[k]; ok {
			filteredRules = append(filteredRules, rule)
		}
	}
	return Analyzer{
		rules:             filteredRules,
		trustedRegistries: trustedRegistries,
	}
}

// Run apply docker best practice rules to docker ast.
func (a Analyzer) Run(node *parser.Node) ([]string, error) {
	var rst []string
	rstChan := make(chan []string, len(a.rules))
	errChan := make(chan error, len(a.rules))

	for i := range a.rules {
		go func(r *rules.Rule) {
			vrst, err := r.ValidateFunc(node, &rules.RuleOptions{
				TrustedRegistries: a.trustedRegistries,
			})
			if err != nil {
				errChan <- err
			} else {
				rstChan <- rules.CreateMessage(a.rules[i], vrst)
			}
		}(a.rules[i])
		select {
		case value := <-rstChan:
			rst = append(rst, value...)
		case err := <-errChan:
			return nil, err
		}
	}
	return rst, nil
}

// getMakeDifference is a function to create a difference set.
func getMakeDiff(xs, ys []string) []string {
	if len(xs) > len(ys) {
		return makeDiff(xs, ys)
	}
	return makeDiff(ys, xs)
}

// make set difference.
func makeDiff(xs, ys []string) []string {
	var set []string
	for i := range xs {
		if !isContain(ys, xs[i]) {
			set = append(set, xs[i])
		}
	}
	return set
}

// isContain is a function to check if s is in xs.
func isContain(xs []string, s string) bool {
	for i := range xs {
		if xs[i] == s {
			return true
		}
	}
	return false
}
