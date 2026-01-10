package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3042 Avoid cache directory with `pip install --no-cache-dir <package>`.
func validateDL3042(node *parser.Node, _ *RuleOptions) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isPip, isInstall, hasNoCacheDir bool
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "pip", "pip3", "pip2":
					isPip = true
					hasNoCacheDir = false
				case "install":
					if isPip {
						isInstall = true
					}
				case "--no-cache-dir":
					if isPip {
						hasNoCacheDir = true
					}
				case "&&":
					if isPip && isInstall && !hasNoCacheDir {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isPip, isInstall, hasNoCacheDir = false, false, false
				}
			}
			if isPip && isInstall && !hasNoCacheDir {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
