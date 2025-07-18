package partials

import (
	"fmt"
	"path"

	"github.com/dhkamp/md-merger/internal/reader"
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
			partialContents[partial] = string(contentBuff)
		}
	}

	return partialContents
}
