package writer

import (
	"fmt"
	"os"
)

func WriteFile(path string, content []byte) error {
	writeErr := os.WriteFile(path, content, 0644)
	if writeErr != nil {
		fmt.Println(writeErr)
		return writeErr
	}
	return nil
}
