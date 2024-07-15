// Package rules provides dockerfile lint rules.
package rules

import (
	"fmt"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Rule is filtered rule (with ignore rule applied)
// ValidateFunc func(node *parser.Node, file string) (rst []string, err error)
type Rule struct {
	Code         string
	Severity     Severity
	Description  string
	ValidateFunc func(node *parser.Node, opts *RuleOptions) ([]ValidateResult, error)
}

type RuleOptions struct {
	TrustedRegistries []string
}

// ValidateResult ValidateFunc's results.
type ValidateResult struct {
	line   int
	addMsg string
}

// Dockerfile instruction.
const (
	FROM       = "FROM"
	LABEL      = "LABEL"
	RUN        = "RUN"
	CMD        = "CMD"
	EXPOSE     = "EXPOSE"
	ADD        = "ADD"
	COPY       = "COPY"
	ENTRYPOINT = "ENTRYPOINT"
	VOLUME     = "VOLUME"
	USER       = "USER"
	WORKDIR    = "WORKDIR"
	SHELL      = "SHELL"

	// deprecated instruction.
	MAINTAINER = "MAINTAINER"
)

// Severity stand check type.
type Severity struct {
	Name string
}

// Severity Level.
var (
	SeverityError      = Severity{Name: "ErrorC"}
	SeverityWarning    = Severity{Name: "WarningC"}
	SeverityInfo       = Severity{Name: "InfoC"}
	SeverityDeprecated = Severity{Name: "Deprecated"}
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
	"DL3018",
	"DL3019",
	"DL3020",
	"DL3021",
	"DL3022",
	"DL3023",
	"DL3024",
	"DL3025",
	"DL3026",
	"DL3027",
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
		Description:  "Use absolute WORKDIR.",
		ValidateFunc: validateDL3000,
	},
	"DL3001": {
		Code:         "DL3001",
		Severity:     SeverityInfo,
		Description:  "For some bash commands it makes no sense running them in a Docker container like `free`, `ifconfig`, `kill`, `mount`, `ps`, `service`, `shutdown`, `ssh`, `top`, `vim`.",
		ValidateFunc: validateDL3001,
	},
	"DL3002": {
		Code:         "DL3002",
		Severity:     SeverityWarning,
		Description:  "Last USER should not be root.",
		ValidateFunc: validateDL3002,
	},
	"DL3003": {
		Code:         "DL3003",
		Severity:     SeverityWarning,
		Description:  "Use WORKDIR to switch to a directory.",
		ValidateFunc: validateDL3003,
	},
	"DL3004": {
		Code:         "DL3004",
		Severity:     SeverityError,
		Description:  "Do not use sudo as it leads to unpredictable behavior. Use a tool like gosu to enforce root.",
		ValidateFunc: validateDL3004,
	},
	"DL3005": {
		Code:         "DL3005",
		Severity:     SeverityError,
		Description:  "Do not use apt-get upgrade or dist-upgrade.",
		ValidateFunc: validateDL3005,
	},
	"DL3006": {
		Code:         "DL3006",
		Severity:     SeverityWarning,
		Description:  "Always tag the version of an image explicitly.",
		ValidateFunc: validateDL3006,
	},
	"DL3007": {
		Code:         "DL3007",
		Severity:     SeverityWarning,
		Description:  "Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.",
		ValidateFunc: validateDL3007,
	},
	"DL3008": {
		Code:         "DL3008",
		Severity:     SeverityWarning,
		Description:  "Pin versions in apt get install. Instead of `apt-get install <package>` use `apt-get install <package>=<version>`.",
		ValidateFunc: validateDL3008,
	},
	"DL3009": {
		Code:         "DL3009",
		Severity:     SeverityInfo,
		Description:  "Delete the apt-get lists after installing something.",
		ValidateFunc: validateDL3009,
	},
	"DL3010": {
		Code:         "DL3010",
		Severity:     SeverityInfo,
		Description:  "Use ADD for extracting archives into an image.",
		ValidateFunc: validateDL3010,
	},
	"DL3011": {
		Code:         "DL3011",
		Severity:     SeverityError,
		Description:  "Valid UNIX ports range from 0 to 65535.",
		ValidateFunc: validateDL3011,
	},
	//"DL3012": {
	//	Code:     "DL3012",
	//	Severity: SeverityDeprecated,
	//	Description:  "Provide an email address or URL as maintainer.",
	//	ValidateFunc:   validateDL3012,
	//},
	"DL3013": {
		Code:         "DL3013",
		Severity:     SeverityWarning,
		Description:  "Pin versions in pip. Instead of `pip install <package>` use `pip install <package>==<version>`.",
		ValidateFunc: validateDL3013,
	},
	"DL3014": {
		Code:         "DL3014",
		Severity:     SeverityWarning,
		Description:  "Use the `-y` switch to avoid manual input `apt-get -y install <package>`.",
		ValidateFunc: validateDL3014,
	},
	"DL3015": {
		Code:         "DL3015",
		Severity:     SeverityInfo,
		Description:  "Avoid additional packages by specifying `--no-install-recommends`.",
		ValidateFunc: validateDL3015,
	},
	"DL3016": {
		Code:         "DL3016",
		Severity:     SeverityWarning,
		Description:  "Pin versions in npm. Instead of `npm install <package>` use `npm install <package>@<version>`.",
		ValidateFunc: validateDL3016,
	},
	"DL3018": {
		Code:         "DL3018",
		Severity:     SeverityWarning,
		Description:  "Pin versions in apk add. Instead of `apk add <package>` use `apk add <package>=<version>`.",
		ValidateFunc: validateDL3018,
	},
	"DL3019": {
		Code:         "DL3019",
		Severity:     SeverityInfo,
		Description:  "Use the `--no-cache` switch to avoid the need to use `--update` and remove `/var/cache/apk/*` when done installing packages.",
		ValidateFunc: validateDL3019,
	},
	"DL3020": {
		Code:         "DL3020",
		Severity:     SeverityError,
		Description:  "Use COPY instead of ADD for files and folders.",
		ValidateFunc: validateDL3020,
	},
	"DL3021": {
		Code:         "DL3021",
		Severity:     SeverityError,
		Description:  "`COPY` with more than 2 arguments requires the last argument to end with `/`.",
		ValidateFunc: validateDL3021,
	},
	"DL3022": {
		Code:         "DL3022",
		Severity:     SeverityWarning,
		Description:  "COPY --from should reference a previously defined FROM alias.",
		ValidateFunc: validateDL3022,
	},
	"DL3023": {
		Code:         "DL3023",
		Severity:     SeverityError,
		Description:  "COPY --from should reference a previously defined FROM alias.",
		ValidateFunc: validateDL3023,
	},
	"DL3024": {
		Code:         "DL3024",
		Severity:     SeverityError,
		Description:  "FROM aliases (stage names) must be unique.",
		ValidateFunc: validateDL3024,
	},
	"DL3025": {
		Code:         "DL3025",
		Severity:     SeverityWarning,
		Description:  "Use arguments JSON notation for CMD and ENTRYPOINT arguments.",
		ValidateFunc: validateDL3025,
	},
	"DL3026": {
		Code:         "DL3026",
		Severity:     SeverityWarning,
		Description:  "Use only an allowed registry in the FROM image.",
		ValidateFunc: validateDL3026,
	},
	"DL3027": {
		Code:         "DL3027",
		Severity:     SeverityWarning,
		Description:  "Do not use apt; use apt-get or apt-cache instead.",
		ValidateFunc: validateDL3027,
	},
	"DL4000": {
		Code:         "DL4000",
		Severity:     SeverityError,
		Description:  "MAINTAINER is deprecated.",
		ValidateFunc: validateDL4000,
	},
	"DL4001": {
		Code:         "DL4001",
		Severity:     SeverityWarning,
		Description:  "Either use Wget or Curl but not both.",
		ValidateFunc: validateDL4001,
	},
	"DL4003": {
		Code:         "DL4003",
		Severity:     SeverityWarning,
		Description:  "Multiple `CMD` instructions found. If you list more than one `CMD` then only the last `CMD` will take effect.",
		ValidateFunc: validateDL4003,
	},
	"DL4004": {
		Code:         "DL4004",
		Severity:     SeverityError,
		Description:  "Multiple `ENTRYPOINT` instructions found. If you list more than one `ENTRYPOINT` then only the last `ENTRYPOINT` will take effect.",
		ValidateFunc: validateDL4004,
	},
	"DL4005": {
		Code:         "DL4005",
		Severity:     SeverityWarning,
		Description:  "Use SHELL to change the default shell.",
		ValidateFunc: validateDL4005,
	},
	"DL4006": {
		Code:         "DL4006",
		Severity:     SeverityWarning,
		Description:  "Set the `SHELL` option -o pipefail before `RUN` with a pipe in it.",
		ValidateFunc: validateDL4006,
	},
}

// isContain is a function to check if s is in xs
func isContain(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

// CreateMessage : create output message
func CreateMessage(rule *Rule, vrst []ValidateResult) []string {
	rst := make([]string, len(vrst))
	for i, v := range vrst {
		rst[i] = fmt.Sprintf("#%d %s %s %s\n", v.line, rule.Code, rule.Description, v.addMsg)
	}
	return rst
}
