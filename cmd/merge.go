package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/dhkamp/md-merger/internal/partials"
	"github.com/dhkamp/md-merger/internal/writer"
	"github.com/spf13/cobra"
)

var input string
var outName string

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge partials",
	Args: func(cmd *cobra.Command, args []string) error {
		_, statErr := os.Stat(input)
		if statErr != nil && errors.Is(statErr, fs.ErrNotExist) {
			return errors.New("input can not be found")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		inputContent, readErr := os.ReadFile(input)
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", readErr)
			os.Exit(1)
		}

		inputDir := path.Dir(input)
		content := string(inputContent)
		if outName == "" {
			outName = getOutFileName(input)
		}

		allUniqPartials := partials.GetAllUniq(&content)
		partialsWithContent := partials.GetContent(allUniqPartials, inputDir)
		mergedContent := partials.Replace(content, partialsWithContent)

		writeErr := writer.WriteFile(path.Join(inputDir, outName+".md"), []byte(mergedContent))
		if writeErr != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", writeErr)
			os.Exit(1)
		}
	},
}

func init() {
	mergeCmd.Flags().StringVarP(&input, "input", "i", "", "Input file used as merge base.")
	mergeCmd.Flags().StringVarP(&outName, "outname", "o", "", "Name of the file containing all the merged partials")

	mergeCmd.MarkFlagRequired("input")

	rootCmd.AddCommand(mergeCmd)
}

func getOutFileName(input string) string {
	base := path.Base(input)
	return base[:len(base)-len(path.Ext(base))] + "_merged"
}
