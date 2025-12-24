package osutil

import (
	"errors"
	"io/fs"
	"os"
)

func CheckExist(fileOrDir string) (bool, error) {
	_, err := os.Stat(fileOrDir)
	if err != nil {
		// file or dir not exist, this err is expected
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}
		// other unexpected error
		return false, err
	}
	return true, nil
}

func CheckDirEmpty(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	return len(entries) == 0, nil
}

func RemoveDirIfExist(dir string) error {
	exist, err := CheckExist(dir)
	if err != nil {
		return err
	}
	// dir not exist, no need to remove
	if !exist {
		return nil
	}

	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return nil
}
