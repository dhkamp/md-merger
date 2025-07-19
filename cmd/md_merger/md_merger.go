package md_merger

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/dhkamp/md-merger/internal/partials"
	"github.com/dhkamp/md-merger/internal/reader"
	"github.com/dhkamp/md-merger/internal/writer"
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

			c, e := filepath.Abs(inputPath)
			if e != nil {
				fmt.Println(e)
				return
			}

			fmt.Println(c)
			fmt.Println(path.Dir(c))

			contentBuff, readErr := reader.ReadFile(inputPath)
			if readErr != nil {
				fmt.Println(readErr)
				return
			}

			dir := path.Dir(inputPath)
			content := string(contentBuff)

			allUniqPartials := partials.GetAll(content)
			partialsWithContent := partials.GetContent(allUniqPartials, dir)
			fmt.Println(partialsWithContent)

			newcontent := partials.Replace(content, partialsWithContent)

			inputFileName := path.Base(inputPath)
			newFileName := inputFileName[:len(inputFileName)-len(path.Ext(inputFileName))]
			writeErr := writer.WriteFile(path.Join(dir, newFileName+"_merged.md"), []byte(newcontent))
			if writeErr != nil {
				fmt.Println(readErr)
				return
			}
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
