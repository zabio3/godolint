package rules

// Rule is filtered rule (with ignore rule applied)
// ValidateFunc func(node *parser.Node, file string) (rst []string, err error)
type Rule struct {
	Code         string
	Severity     Severity
	ValidateFunc interface{}
}

// Severity stand check type
type Severity struct {
	Name string
}

// Severity Level
var (
	SeverityError   = Severity{Name: "ErrorC"}
	SeverityWarning = Severity{Name: "WarningC"}
	SeverityInfo    = Severity{Name: "InfoC"}
	//SeverityDeprecated = Severity{Name: "Deprecated"}
)

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
}

// Rules (Docker best practice rule key)
var Rules = map[string]*Rule{
	"DL3000": {
		Code:         "DL3000",
		Severity:     SeverityError,
		ValidateFunc: validateDL3000,
	},
	"DL3001": {
		Code:         "DL3001",
		Severity:     SeverityInfo,
		ValidateFunc: validateDL3001,
	},
	"DL3002": {
		Code:         "DL3002",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3002,
	},
	"DL3003": {
		Code:         "DL3003",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3003,
	},
	"DL3004": {
		Code:         "DL3004",
		Severity:     SeverityError,
		ValidateFunc: validateDL3004,
	},
	"DL3005": {
		Code:         "DL3005",
		Severity:     SeverityError,
		ValidateFunc: validateDL3005,
	},
	"DL3006": {
		Code:         "DL3006",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3006,
	},
	"DL3007": {
		Code:         "DL3007",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3007,
	},
	"DL3008": {
		Code:         "DL3008",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3008,
	},
	"DL3009": {
		Code:         "DL3009",
		Severity:     SeverityInfo,
		ValidateFunc: validateDL3009,
	},
	"DL3010": {
		Code:         "DL3010",
		Severity:     SeverityInfo,
		ValidateFunc: validateDL3010,
	},
	"DL3011": {
		Code:         "DL3011",
		Severity:     SeverityError,
		ValidateFunc: validateDL3011,
	},
	//"DL3012": {
	//	Code:     "DL3012",
	//	Severity: SeverityDeprecated,
	//	ValidateFunc:   validateDL3012,
	//},
	"DL3013": {
		Code:         "DL3013",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3013,
	},
	"DL3014": {
		Code:         "DL3014",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3014,
	},
	"DL3015": {
		Code:         "DL3015",
		Severity:     SeverityInfo,
		ValidateFunc: validateDL3015,
	},
	"DL3016": {
		Code:         "DL3016",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3016,
	},
	"DL3017": {
		Code:         "DL3017",
		Severity:     SeverityError,
		ValidateFunc: validateDL3017,
	},
	"DL3018": {
		Code:         "DL3018",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3018,
	},
	"DL3019": {
		Code:         "DL3019",
		Severity:     SeverityInfo,
		ValidateFunc: validateDL3019,
	},
	"DL3020": {
		Code:         "DL3020",
		Severity:     SeverityError,
		ValidateFunc: validateDL3020,
	},
	"DL3021": {
		Code:         "DL3021",
		Severity:     SeverityError,
		ValidateFunc: validateDL3021,
	},
	"DL3022": {
		Code:         "DL3022",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3022,
	},
	"DL3023": {
		Code:         "DL3023",
		Severity:     SeverityError,
		ValidateFunc: validateDL3023,
	},
	"DL3024": {
		Code:         "DL3024",
		Severity:     SeverityError,
		ValidateFunc: validateDL3024,
	},
	"DL3025": {
		Code:         "DL3025",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL3025,
	},
	"DL4000": {
		Code:         "DL4000",
		Severity:     SeverityError,
		ValidateFunc: validateDL4000,
	},
	"DL4001": {
		Code:         "DL4001",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL4001,
	},
	"DL4003": {
		Code:         "DL4003",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL4003,
	},
	"DL4004": {
		Code:         "DL4004",
		Severity:     SeverityError,
		ValidateFunc: validateDL4004,
	},
	"DL4005": {
		Code:         "DL4005",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL4005,
	},
	"DL4006": {
		Code:         "DL4006",
		Severity:     SeverityWarning,
		ValidateFunc: validateDL4006,
	},
}

func isContains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
