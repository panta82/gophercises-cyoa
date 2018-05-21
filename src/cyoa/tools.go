package cyoa

import (
	"path/filepath"
	"os"
)

func getAssetPath(fileName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("Failed to determine directory of the executable")
	}

	return filepath.Join(dir, fileName)
}
