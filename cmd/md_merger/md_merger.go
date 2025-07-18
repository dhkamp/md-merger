package md_merger

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var (
	inputPath string
	outDir    string

	rootCmd = &cobra.Command{
		Use:   "md-merger",
		Short: "md-merger is a CLI tool to merge markdown files",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(inputPath) == 0 {
				return errors.New("missing required flag --input")
			}

			_, statErr := os.Stat(inputPath)
			if statErr != nil && errors.Is(statErr, fs.ErrNotExist) {
				return errors.New("input file does not exist")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(outDir)
			contentBuff, readErr := os.ReadFile(inputPath)
			if readErr != nil {
				fmt.Println(readErr)
				return
			}

			dir := path.Dir(inputPath)

			content := string(contentBuff)
			partials := findAllPartials(content)

			partialsContents := getPartialContent(dir, partials)
			newcontent := replacePartials(content, partialsContents)

			fmt.Println(newcontent)

			inputFileName := path.Base(inputPath)
			newFileName := inputFileName[:len(inputFileName)-len(path.Ext(inputFileName))]
			writeErr := os.WriteFile(path.Join(dir, newFileName+"_merged.md"), []byte(newcontent), 0644)
			if writeErr != nil {
				fmt.Println(readErr)
				return
			}
		},
	}
)

func findAllPartials(content string) []string {
	regex := regexp.MustCompile(`<!-- merge:.* -->`)
	matches := regex.FindAllStringSubmatchIndex(content, -1)
	var partials []string

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		partial := content[match[0]:match[1]]
		partialPath := strings.Replace(strings.Replace(partial, "<!-- merge:", "", 1), " -->", "", 1)

		if !slices.Contains(partials, partialPath) {
			partials = append(partials, partialPath)
		}
	}
	return partials
}

func getPartialContent(basePath string, partials []string) map[string]string {

	partialContents := make(map[string]string)

	for i := 0; i < len(partials); i++ {
		partial := partials[i]
		partialPath := path.Join(basePath, partial)
		contentBuff, readErr := os.ReadFile(partialPath)
		if readErr != nil {
			fmt.Println(readErr)
		} else {
			partialContents[partial] = string(contentBuff)
		}
	}

	return partialContents
}

func replacePartials(baseContent string, partials map[string]string) string {
	newContent := baseContent
	for key, val := range partials {
		newContent = strings.ReplaceAll(newContent, "<!-- merge:"+key+" -->", val)
	}
	return newContent
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&inputPath, "input", "", "Output directory")
	rootCmd.PersistentFlags().StringVar(&outDir, "out", "", "Output directory")
}
