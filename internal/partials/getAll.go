package partials

import (
	"regexp"
	"slices"
	"strings"
)

func GetAllUniq(content *string) []string {
	partialsRx := regexp.MustCompile(`<!-- merge:.* -->`)
	matches := partialsRx.FindAllStringSubmatchIndex(*content, -1)
	var partials []string

	if matches == nil {
		return partials
	}

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		partial := (*content)[match[0]:match[1]]
		partialPath := strings.Replace(strings.Replace(partial, "<!-- merge:", "", 1), " -->", "", 1)
		if !slices.Contains(partials, partialPath) {
			partials = append(partials, partialPath)
		}
	}

	return partials
}
