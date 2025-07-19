package partials

import (
	"fmt"
	"path"
	"regexp"
	"slices"
	"strings"

	"github.com/dhkamp/md-merger/internal/reader"
)

var (
	relativeLinksRx *regexp.Regexp = regexp.MustCompile(`\[([^\]]*?)\]\(([^)]+?)\)`)
)

func GetContent(partials []string, baseDir string) map[string]string {
	partialContents := make(map[string]string)

	for i := 0; i < len(partials); i++ {
		partial := partials[i]
		partialPath := path.Join(baseDir, partial)
		contentBuff, err := reader.ReadFile(partialPath)
		if err != nil {
			// TODO: handle
			fmt.Println(err)
		} else {
			mdcontent := updateMDContents(string(contentBuff), path.Dir(partial))
			partialContents[partial] = mdcontent
		}
	}

	return partialContents
}

func updateMDContents(content string, baseDir string) string {
	var replaced []string
	newContent := content
	matches := relativeLinksRx.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) >= 3 && !slices.Contains(replaced, match[2]) {
			newContent = strings.ReplaceAll(newContent, match[2], path.Join(baseDir, match[2]))
			replaced = append(replaced, match[2])
		}
	}

	return newContent
}
