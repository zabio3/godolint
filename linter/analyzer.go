package linter

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/zabio3/godolint/linter/rules"
)

// Analyzer Apply docker best practice rules to docker ast
func Analyzer(node *parser.Node, file string, ignoreRules []string) ([]string, error) {
	var (
		rst           []string
		filteredRules []string
	)

	// Filter rules to apply
	if len(ignoreRules) != 0 {
		for _, v := range ignoreRules {
			xs, err := getFilterdList(v, rules.RuleKeys)
			if err != nil {
				return nil, err
			}
			filteredRules = xs
		}
	} else {
		filteredRules = rules.RuleKeys
	}

	for _, k := range filteredRules {
		v, err := rules.Rules[k].CheckF.(func(node *parser.Node, file string) (rst []string, err error))(node, file)
		if err != nil {
			return rst, err
		}
		rst = append(rst, v...)
	}
	return rst, nil
}

func getFilterdList(s string, xs []string) ([]string, error) {
	var filteredRules []string
	isExist := false
	for _, x := range xs {
		if x == s {
			isExist = true
		} else {
			filteredRules = append(filteredRules, x)
		}
	}

	if !isExist {
		return nil, fmt.Errorf("no exist rule specified by ignore flag: %s", s)
	}

	return filteredRules, nil
}