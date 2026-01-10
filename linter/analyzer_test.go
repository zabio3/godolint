package linter

import (
	"slices"
	"testing"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func TestAnalyzer(t *testing.T) {
	cases := []struct {
		name        string
		node        *parser.Node
		ignoreRules []string
		expectedRst []string
		expectedErr error
	}{
		{
			name: "relative WORKDIR violation",
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
			name: "all rules ignored",
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

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			analyzer := NewAnalyzer(tc.ignoreRules, nil)
			gotRst, gotErr := analyzer.Run(tc.node)
			if !slices.Equal(tc.expectedRst, gotRst) {
				t.Errorf("results deep equal has returned: want %s, got %s", tc.expectedRst, gotRst)
			}

			if gotErr != tc.expectedErr {
				t.Errorf("unexpected outStream has returned: want: %s, got: %s", tc.expectedErr, gotErr)
			}
		})
	}
}
