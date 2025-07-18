package partials

import "strings"

func Replace(content string, partials map[string]string) string {
	newContent := content
	for key, val := range partials {
		newContent = strings.ReplaceAll(newContent, "<!-- merge:"+key+" -->", val)
	}
	return newContent
}
