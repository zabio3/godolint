package linter

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Analize Apply docker best practice rules to docker ast
func Analize(node *parser.Node, file string, ignoreRules []string) ([]string, error) {
	var (
		rst           []string
		filteredRules []string
	)

	// Filtering rules to apply
	if len(ignoreRules) != 0 {
		for _, v := range ignoreRules {
			for _, w := range RuleKeys {
				if v != w {
					filteredRules = append(filteredRules, w)
				}
			}
		}
	} else {
		filteredRules = RuleKeys
	}

	for _, k := range filteredRules {
		v, err := Rules[k].CheckF.(func(node *parser.Node, file string) (rst []string, err error))(node, file)
		if err != nil {
			return rst, err
		}
		for _, w := range v {
			rst = append(rst, w)
		}
	}
	return rst, nil
}
