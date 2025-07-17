package md_merger

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

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
			content := string(contentBuff)
			fmt.Println(content)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&inputPath, "input", "", "Output directory")
	rootCmd.PersistentFlags().StringVar(&outDir, "out", "", "Output directory")
}
