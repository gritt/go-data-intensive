package tags

import (
	"regexp"
	"strings"
)

func ParseTitle(html string) string {
	parser, err := regexp.Compile(`<title>\s*(.*?)\s*</title>`)
	if err != nil {
		// TODO log err: failed to compile title exp
		return ""
	}

	data := parser.FindSubmatch([]byte(html))
	if len(data) < 2 {
		// TODO log err: title not found
		return ""
	}

	title := strings.TrimSpace(string(data[1]))

	return title
}
