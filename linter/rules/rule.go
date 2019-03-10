package rules

// Rule is filtered rule (with ignore rule applied)
// CheckF func(node *parser.Node, file string) (rst []string, err error)
type Rule struct {
	Code     string
	Severity string
	CheckF   interface{}
}

// RuleKeys is (Docker best practice rule key)
var RuleKeys = []string{
	"DL3000",
	"DL3001",
	"DL3002",
	"DL3003",
	"DL3004",
	"DL3005",
}

// Rules (Docker best practice rule key)
var Rules = map[string]*Rule{
	"DL3000": {
		Code:     "DL3000",
		Severity: "ErrorC",
		CheckF:   dl3000Check,
	},
	"DL3001": {
		Code:     "DL3001",
		Severity: "InfoC",
		CheckF:   dl3001Check,
	},
	"DL3002": {
		Code:     "DL3002",
		Severity: "WarningC",
		CheckF:   dl3002Check,
	},
	"DL3003": {
		Code:     "DL3003",
		Severity: "WarningC",
		CheckF:   dl3003Check,
	},
	"DL3004": {
		Code:     "DL3004",
		Severity: "ErrorC",
		CheckF:   dl3004Check,
	},
	"DL3005": {
		Code:     "DL3005",
		Severity: "ErrorC",
		CheckF:   dl3005Check,
	},
}
