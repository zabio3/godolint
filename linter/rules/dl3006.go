package rules

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexDL3006 = regexp.MustCompile(`.+[:].+`)

// validateDL3006 Always tag the version of an image explicitly"
func validateDL3006(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == FROM {
			inst, parseErr := instructions.ParseInstruction(child)
			st, ok := inst.(*instructions.Stage)
			if parseErr != nil || !ok {
				continue
			}

			if st.BaseName == "scratch" {
				continue
			}

			if !regexDL3006.MatchString(child.Next.Value) {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
