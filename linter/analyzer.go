// Package linter provides dockerfile analyzer (Apply the rule to the parsed dockerfile).
package linter

import (
	"slices"

	"github.com/moby/buildkit/frontend/dockerfile/parser"

	"github.com/zabio3/godolint/linter/rules"
)

// Analyzer implements Analyzer.
type Analyzer struct {
	rules             []*rules.Rule
	trustedRegistries []string
}

// NewAnalyzer generates a new Analyzer with rules to apply.
func NewAnalyzer(ignoreRules []string, trustedRegistries []string) Analyzer {
	var filteredRules []*rules.Rule
	for _, key := range rules.RuleKeys {
		if slices.Contains(ignoreRules, key) {
			continue
		}
		if rule, ok := rules.Rules[key]; ok {
			filteredRules = append(filteredRules, rule)
		}
	}
	return Analyzer{
		rules:             filteredRules,
		trustedRegistries: trustedRegistries,
	}
}

// Run applies docker best practice rules to docker ast.
func (a Analyzer) Run(node *parser.Node) ([]string, error) {
	var rst []string
	for _, rule := range a.rules {
		vrst, err := rule.ValidateFunc(node, &rules.RuleOptions{
			TrustedRegistries: a.trustedRegistries,
		})
		if err != nil {
			return nil, err
		}
		rst = append(rst, rules.CreateMessage(rule, vrst)...)
	}
	return rst, nil
}
