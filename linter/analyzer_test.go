package linter

import (
	"testing"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func TestAnalyzer(t *testing.T) {
	cases := []struct {
		node        *parser.Node
		ignoreRules []string
		expectedRst []string
		expectedErr error
	}{
		{
			node: &parser.Node{
				Value: "",
				Next:  (*parser.Node)(nil),
				Children: []*parser.Node{
					{
						Value: "FROM",
						Next: &parser.Node{
							Value:      "golang:1.12.0-stretch",
							Next:       (*parser.Node)(nil),
							Children:   nil,
							Attributes: nil,
							Original:   "",
							Flags:      nil,
							StartLine:  0,
						},
						Children:   nil,
						Attributes: nil,
						Original:   "FROM golang:1.12.0-stretch",
						Flags:      nil,
						StartLine:  1,
					},
					{
						Value: "WORKDIR",
						Next: &parser.Node{
							Value:      "go/",
							Next:       (*parser.Node)(nil),
							Children:   nil,
							Attributes: nil,
							Original:   "",
							Flags:      nil,
							StartLine:  0,
						},
						Children:   nil,
						Attributes: nil,
						Original:   "WORKDIR go/",
						Flags:      nil,
						StartLine:  3,
					},
					{
						Value: "COPY",
						Next: &parser.Node{
							Value: ".",
							Next: &parser.Node{
								Value:      "/go",
								Next:       (*parser.Node)(nil),
								Children:   nil,
								Attributes: nil,
								Original:   "",
								Flags:      nil,
								StartLine:  0,
							},
							Children:   nil,
							Attributes: nil,
							Original:   "",
							Flags:      nil,
							StartLine:  0,
						},
						Children:   nil,
						Attributes: nil,
						Original:   "COPY . /go",
						Flags:      nil,
						StartLine:  4,
					},
					{
						Value: "CMD",
						Next: &parser.Node{
							Value: "go",
							Next: &parser.Node{
								Value: "run",
								Next: &parser.Node{
									Value:      "main.go",
									Next:       (*parser.Node)(nil),
									Children:   nil,
									Attributes: nil,
									Original:   "",
									Flags:      nil,
									StartLine:  0,
								},
								Children:   nil,
								Attributes: nil,
								Original:   "",
								Flags:      nil,
								StartLine:  0,
							},
							Children:   nil,
							Attributes: nil,
							Original:   "",
							Flags:      nil,
							StartLine:  0,
						},
						Children: nil,
						Attributes: map[string]bool{
							"json": true,
						},
						Original:  "CMD [\"go\", \"run\", \"main.go\"]",
						Flags:     nil,
						StartLine: 6,
					},
				},
				Attributes: nil,
				Original:   "",
				Flags:      nil,
				StartLine:  1,
			},
			ignoreRules: []string{
				"DL4000",
			},
			expectedRst: []string{
				"#3 DL3000 Use absolute WORKDIR. \n",
			},
			expectedErr: nil,
		},
		{
			node: nil,
			ignoreRules: []string{
				"DL3000",
				"DL3001",
				"DL3002",
				"DL3003",
				"DL3004",
				"DL3005",
				"DL3006",
				"DL3007",
				"DL3008",
				"DL3009",
				"DL3010",
				"DL3011",
				//"DL3012",
				"DL3013",
				"DL3014",
				"DL3015",
				"DL3016",
				"DL3018",
				"DL3019",
				"DL3020",
				"DL3021",
				"DL3022",
				"DL3023",
				"DL3024",
				"DL3025",
				"DL3027",
				"DL4000",
				"DL4001",
				"DL4003",
				"DL4004",
				"DL4005",
				"DL4006",
			},
			expectedRst: nil,
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		analyzer := NewAnalyzer(tc.ignoreRules, nil)
		gotRst, gotErr := analyzer.Run(tc.node)
		if !sliceEq(tc.expectedRst, gotRst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d Unexpected outStream has returned: want: %s, got: %s", i, tc.expectedErr, gotErr)
		}

		cleanup(t)
	}

}

func TestGetMakeDifference(t *testing.T) {
	cases := []struct {
		ignoreRules []string
		rules       []string
		expectedRst []string
	}{
		{
			ignoreRules: []string{
				"DL3001",
				"DL3018",
				"DL4004",
			},
			rules: []string{
				"DL3000",
				"DL3001",
				"DL3002",
				"DL3003",
				"DL3004",
				"DL3005",
				"DL3006",
				"DL3007",
				"DL3008",
				"DL3009",
				"DL3010",
				"DL3011",
				//"DL3012",
				"DL3013",
				"DL3014",
				"DL3015",
				"DL3016",
				"DL3018",
				"DL3019",
				"DL3020",
				"DL3021",
				"DL3022",
				"DL3023",
				"DL3024",
				"DL3025",
				"DL4000",
				"DL4001",
				"DL4003",
				"DL4004",
				"DL4005",
				"DL4006",
			},
			expectedRst: []string{
				"DL3000",
				"DL3002",
				"DL3003",
				"DL3004",
				"DL3005",
				"DL3006",
				"DL3007",
				"DL3008",
				"DL3009",
				"DL3010",
				"DL3011",
				//"DL3012",
				"DL3013",
				"DL3014",
				"DL3015",
				"DL3016",
				"DL3019",
				"DL3020",
				"DL3021",
				"DL3022",
				"DL3023",
				"DL3024",
				"DL3025",
				"DL4000",
				"DL4001",
				"DL4003",
				"DL4005",
				"DL4006",
			},
		},
		{
			ignoreRules: []string{
				"DL3000",
				"DL3001",
				"DL3002",
				"DL3003",
				"DL3004",
				"DL3005",
				"DL3006",
				"DL3007",
				"DL3008",
				"DL3009",
				"DL3010",
				"DL3011",
				//"DL3012",
				"DL3013",
				"DL3014",
				"DL3015",
				"DL3016",
				"DL3018",
				"DL3019",
				"DL3020",
				"DL3021",
				"DL3022",
				"DL3023",
				"DL3024",
				"DL3025",
				"DL4000",
				"DL4001",
				"DL4003",
				"DL4004",
				"DL4005",
				"DL4006",
			},
			rules: []string{
				"DL3001",
				"DL3018",
				"DL4004",
			},
			expectedRst: []string{
				"DL3000",
				"DL3002",
				"DL3003",
				"DL3004",
				"DL3005",
				"DL3006",
				"DL3007",
				"DL3008",
				"DL3009",
				"DL3010",
				"DL3011",
				//"DL3012",
				"DL3013",
				"DL3014",
				"DL3015",
				"DL3016",
				"DL3019",
				"DL3020",
				"DL3021",
				"DL3022",
				"DL3023",
				"DL3024",
				"DL3025",
				"DL4000",
				"DL4001",
				"DL4003",
				"DL4005",
				"DL4006",
			},
		},
	}

	for i, tc := range cases {
		rst := getMakeDiff(tc.rules, tc.ignoreRules)
		if !sliceEq(tc.expectedRst, rst) {
			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, rst)
		}
		cleanup(t)
	}
}

func cleanup(t *testing.T) {
	t.Helper()
}

// reflect.DeepEqual(gotRst, gotRst)
func sliceEq(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
