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
	FROM        = "FROM"
	LABEL       = "LABEL"
	RUN         = "RUN"
	CMD         = "CMD"
	EXPOSE      = "EXPOSE"
	ADD         = "ADD"
	COPY        = "COPY"
	ENTRYPOINT  = "ENTRYPOINT"
	VOLUME      = "VOLUME"
	USER        = "USER"
	WORKDIR     = "WORKDIR"
	SHELL       = "SHELL"
	HEALTHCHECK = "HEALTHCHECK"
	ONBUILD     = "ONBUILD"
	ARG         = "ARG"
	ENV         = "ENV"

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
	SeverityStyle      = Severity{Name: "Style"}
	SeverityIgnore     = Severity{Name: "Ignore"}
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
	"DL3012",
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
	"DL3028",
	"DL3029",
	"DL3030",
	"DL3032",
	"DL3033",
	"DL3034",
	"DL3035",
	"DL3036",
	"DL3037",
	"DL3038",
	"DL3040",
	"DL3041",
	"DL3042",
	"DL3043",
	"DL3044",
	"DL3045",
	"DL3046",
	"DL3047",
	"DL3048",
	"DL3049",
	"DL3050",
	"DL3051",
	"DL3052",
	"DL3053",
	"DL3054",
	"DL3055",
	"DL3056",
	"DL3057",
	"DL3058",
	"DL3059",
	"DL3060",
	"DL3061",
	"DL3062",
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
	"DL3012": {
		Code:         "DL3012",
		Severity:     SeverityError,
		Description:  "Multiple `HEALTHCHECK` instructions found. Only the last `HEALTHCHECK` will take effect.",
		ValidateFunc: validateDL3012,
	},
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
	"DL3028": {
		Code:         "DL3028",
		Severity:     SeverityWarning,
		Description:  "Pin versions in gem install. Instead of `gem install <package>` use `gem install <package>:<version>`.",
		ValidateFunc: validateDL3028,
	},
	"DL3029": {
		Code:         "DL3029",
		Severity:     SeverityWarning,
		Description:  "Do not use --platform flag with FROM.",
		ValidateFunc: validateDL3029,
	},
	"DL3030": {
		Code:         "DL3030",
		Severity:     SeverityWarning,
		Description:  "Use the `-y` switch to avoid manual input `yum install -y <package>`.",
		ValidateFunc: validateDL3030,
	},
	"DL3032": {
		Code:         "DL3032",
		Severity:     SeverityWarning,
		Description:  "`yum clean all` missing after yum command.",
		ValidateFunc: validateDL3032,
	},
	"DL3033": {
		Code:         "DL3033",
		Severity:     SeverityWarning,
		Description:  "Specify version with `yum install -y <package>-<version>`.",
		ValidateFunc: validateDL3033,
	},
	"DL3034": {
		Code:         "DL3034",
		Severity:     SeverityWarning,
		Description:  "Non-interactive switch missing from `zypper` command. Use `zypper -n`.",
		ValidateFunc: validateDL3034,
	},
	"DL3035": {
		Code:         "DL3035",
		Severity:     SeverityWarning,
		Description:  "Do not use `zypper dist-upgrade`.",
		ValidateFunc: validateDL3035,
	},
	"DL3036": {
		Code:         "DL3036",
		Severity:     SeverityWarning,
		Description:  "`zypper clean` missing after zypper use.",
		ValidateFunc: validateDL3036,
	},
	"DL3037": {
		Code:         "DL3037",
		Severity:     SeverityWarning,
		Description:  "Specify version with `zypper install <package>=<version>`.",
		ValidateFunc: validateDL3037,
	},
	"DL3038": {
		Code:         "DL3038",
		Severity:     SeverityWarning,
		Description:  "Use the `-y` switch to avoid manual input `dnf install -y <package>`.",
		ValidateFunc: validateDL3038,
	},
	"DL3040": {
		Code:         "DL3040",
		Severity:     SeverityWarning,
		Description:  "`dnf clean all` missing after dnf command.",
		ValidateFunc: validateDL3040,
	},
	"DL3041": {
		Code:         "DL3041",
		Severity:     SeverityWarning,
		Description:  "Specify version with `dnf install -y <package>-<version>`.",
		ValidateFunc: validateDL3041,
	},
	"DL3042": {
		Code:         "DL3042",
		Severity:     SeverityWarning,
		Description:  "Avoid cache directory with `pip install --no-cache-dir <package>`.",
		ValidateFunc: validateDL3042,
	},
	"DL3043": {
		Code:         "DL3043",
		Severity:     SeverityError,
		Description:  "`ONBUILD`, `FROM` or `MAINTAINER` triggered from within `ONBUILD` instruction.",
		ValidateFunc: validateDL3043,
	},
	"DL3044": {
		Code:         "DL3044",
		Severity:     SeverityError,
		Description:  "Do not refer to an environment variable within the same ENV statement where it is defined.",
		ValidateFunc: validateDL3044,
	},
	"DL3045": {
		Code:         "DL3045",
		Severity:     SeverityWarning,
		Description:  "`COPY` to a relative destination without `WORKDIR` set.",
		ValidateFunc: validateDL3045,
	},
	"DL3046": {
		Code:         "DL3046",
		Severity:     SeverityWarning,
		Description:  "`useradd` without flag `-l` and target UID greater than or equal to 65534 can lead to excessively large Image.",
		ValidateFunc: validateDL3046,
	},
	"DL3047": {
		Code:         "DL3047",
		Severity:     SeverityInfo,
		Description:  "`wget` without flag `--progress` will result in excessively bloated build logs when downloading larger files.",
		ValidateFunc: validateDL3047,
	},
	"DL3048": {
		Code:         "DL3048",
		Severity:     SeverityStyle,
		Description:  "Invalid label key.",
		ValidateFunc: validateDL3048,
	},
	"DL3049": {
		Code:         "DL3049",
		Severity:     SeverityIgnore,
		Description:  "Label is missing.",
		ValidateFunc: validateDL3049,
	},
	"DL3050": {
		Code:         "DL3050",
		Severity:     SeverityIgnore,
		Description:  "Superfluous label(s) present.",
		ValidateFunc: validateDL3050,
	},
	"DL3051": {
		Code:         "DL3051",
		Severity:     SeverityWarning,
		Description:  "Label is empty.",
		ValidateFunc: validateDL3051,
	},
	"DL3052": {
		Code:         "DL3052",
		Severity:     SeverityWarning,
		Description:  "Label is not a valid URL.",
		ValidateFunc: validateDL3052,
	},
	"DL3053": {
		Code:         "DL3053",
		Severity:     SeverityWarning,
		Description:  "Label is not a valid RFC3339 format datetime.",
		ValidateFunc: validateDL3053,
	},
	"DL3054": {
		Code:         "DL3054",
		Severity:     SeverityWarning,
		Description:  "Label is not a valid SPDX license identifier.",
		ValidateFunc: validateDL3054,
	},
	"DL3055": {
		Code:         "DL3055",
		Severity:     SeverityWarning,
		Description:  "Label is not a valid git hash.",
		ValidateFunc: validateDL3055,
	},
	"DL3056": {
		Code:         "DL3056",
		Severity:     SeverityWarning,
		Description:  "Label does not conform to semantic versioning.",
		ValidateFunc: validateDL3056,
	},
	"DL3057": {
		Code:         "DL3057",
		Severity:     SeverityIgnore,
		Description:  "`HEALTHCHECK` instruction missing.",
		ValidateFunc: validateDL3057,
	},
	"DL3058": {
		Code:         "DL3058",
		Severity:     SeverityWarning,
		Description:  "Label is not a valid RFC5322 email format.",
		ValidateFunc: validateDL3058,
	},
	"DL3059": {
		Code:         "DL3059",
		Severity:     SeverityInfo,
		Description:  "Multiple consecutive `RUN` instructions. Consider consolidation.",
		ValidateFunc: validateDL3059,
	},
	"DL3060": {
		Code:         "DL3060",
		Severity:     SeverityInfo,
		Description:  "`yarn cache clean` missing after `yarn install`.",
		ValidateFunc: validateDL3060,
	},
	"DL3061": {
		Code:         "DL3061",
		Severity:     SeverityError,
		Description:  "Invalid instruction order. Dockerfile must begin with `FROM`, `ARG` or comment.",
		ValidateFunc: validateDL3061,
	},
	"DL3062": {
		Code:         "DL3062",
		Severity:     SeverityWarning,
		Description:  "Pin versions in go install. Instead of `go install <package>` use `go install <package>@<version>`.",
		ValidateFunc: validateDL3062,
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

// CreateMessage creates output messages from validation results.
func CreateMessage(rule *Rule, vrst []ValidateResult) []string {
	rst := make([]string, len(vrst))
	for i, v := range vrst {
		rst[i] = fmt.Sprintf("#%d %s %s %s\n", v.line, rule.Code, rule.Description, v.addMsg)
	}
	return rst
}
