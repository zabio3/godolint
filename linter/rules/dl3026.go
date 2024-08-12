package rules

import (
	"regexp"
	"strings"

	"github.com/distribution/reference"
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var defaultAllowedRegistries = []string{
	"docker.io",
	"hub.docker.com",
}

// validateDL3026 Use only an allowed registry in the FROM image
func validateDL3026(node *parser.Node, opts *RuleOptions) (rst []ValidateResult, err error) {
	if len(opts.TrustedRegistries) == 0 {
		return rst, nil
	}

	tr := append(opts.TrustedRegistries, defaultAllowedRegistries...)
	for _, child := range node.Children {
		if child.Value == FROM {

			inst, err := instructions.ParseInstruction(child)
			st, ok := inst.(*instructions.Stage)
			if err != nil || !ok {
				// Malformed FROM command
				continue
			}

			if st.BaseName == "scratch" {
				continue
			}

			ref, err := reference.ParseNormalizedNamed(st.BaseName)
			if err != nil {
				continue
			}

			// NOTE: if domain is implicitly docker.io, eg "FROM ubuntu:18.04", this reference.Domain() returns docker.io
			dm := reference.Domain(ref)
			if !isRegistryAllowed(dm, tr) {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}

	return rst, nil
}

// Handle wildcards by converting to regexp
func wildcardToRegexp(url string) string {
	var rb strings.Builder
	components := strings.Split(url, "*")
	if len(components) == 1 {
		return "^" + url + "$" // Exact match
	}

	for i, literal := range components {
		if i > 0 {
			rb.WriteString(".*")
		}

		rb.WriteString(regexp.QuoteMeta(literal))
	}
	return rb.String()
}

func isRegistryAllowed(url string, allowedRegistries []string) bool {
	for _, allowedRegistry := range allowedRegistries {
		re := wildcardToRegexp(allowedRegistry)
		result, err := regexp.MatchString(re, url)
		if err == nil && result {
			return true
		}
	}
	return false
}
