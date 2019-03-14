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
	"DL3017",
	"DL3018",
	"DL3019",
	"DL3020",
	"DL3021",
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
	"DL3006": {
		Code:     "DL3006",
		Severity: "WarningC",
		CheckF:   dl3006Check,
	},
	"DL3007": {
		Code:     "DL3007",
		Severity: "WarningC",
		CheckF:   dl3007Check,
	},
	"DL3008": {
		Code:     "DL3008",
		Severity: "WarningC",
		CheckF:   dl3008Check,
	},
	"DL3009": {
		Code:     "DL3009",
		Severity: "InfoC",
		CheckF:   dl3009Check,
	},
	"DL3010": {
		Code:     "DL3010",
		Severity: "InfoC",
		CheckF:   dl3010Check,
	},
	"DL3011": {
		Code:     "DL3011",
		Severity: "ErrorC",
		CheckF:   dl3011Check,
	},
	//"DL3012": {
	//	Code:     "DL3012",
	//	Severity: "Deprecated",
	//	CheckF:   dl3012Check,
	//},
	"DL3013": {
		Code:     "DL3013",
		Severity: "WarningC",
		CheckF:   dl3013Check,
	},
	"DL3014": {
		Code:     "DL3014",
		Severity: "WarningC",
		CheckF:   dl3014Check,
	},
	"DL3015": {
		Code:     "DL3015",
		Severity: "InfoC",
		CheckF:   dl3015Check,
	},
	"DL3016": {
		Code:     "DL3016",
		Severity: "WarningC",
		CheckF:   dl3016Check,
	},
	"DL3017": {
		Code:     "DL3017",
		Severity: "ErrorC",
		CheckF:   dl3017Check,
	},
	"DL3018": {
		Code:     "DL3018",
		Severity: "WarningC",
		CheckF:   dl3018Check,
	},
	"DL3019": {
		Code:     "DL3019",
		Severity: "InfoC",
		CheckF:   dl3019Check,
	},
	"DL3020": {
		Code:     "DL3020",
		Severity: "ErrorC",
		CheckF:   dl3020Check,
	},
	"DL3021": {
		Code:     "DL3021",
		Severity: "ErrorC",
		CheckF:   dl3021Check,
	},
}
