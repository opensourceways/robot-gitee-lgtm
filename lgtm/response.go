package lgtm

import (
	"fmt"
	"strings"
)

// AboutThisBotWithoutCommands contains the message that explains how to interact with the bot.
const AboutThisBotWithoutCommands = "Instructions for interacting with me using PR comments are available [here](https://git.k8s.io/community/contributors/guide/pull-requests.md).  If you have questions or suggestions related to my behavior, please file an issue against the [opensourceways/test-infra](https://github.com/opensourceways/test-infra/issues/new?title=Prow%20issue:) repository."

// FormatResponse nicely formats a response to a generic reason.
func FormatResponse(to, message, reason string) string {
	format := `@%s: %s

<details>

%s

%s
</details>`

	return fmt.Sprintf(format, to, message, reason, AboutThisBotWithoutCommands)
}

// FormatResponseRaw nicely formats a response for one does not have an issue comment
func FormatResponseRaw(body, bodyURL, login, reply string) string {
	format := `In response to [this](%s):

%s
`
	// Quote the user's comment by prepending ">" to each line.
	var quoted []string
	for _, l := range strings.Split(body, "\n") {
		quoted = append(quoted, ">"+l)
	}
	return FormatResponse(login, reply, fmt.Sprintf(format, bodyURL, strings.Join(quoted, "\n")))
}
