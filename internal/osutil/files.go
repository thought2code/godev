package osutil

import (
	"fmt"
	"os"
	"path/filepath"
)

func ClearDir(dir string) error {
	if dir == "" || dir == "/" || dir == "." || dir == ".." {
		return fmt.Errorf("dir %s is not allowed to clear", dir)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}

	return nil
}
