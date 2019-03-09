package linter

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func Analize(node *parser.Node, file string) ([]string, error) {
	var rst []string
	for _, k := range RuleKeys {
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
