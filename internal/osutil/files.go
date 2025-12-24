package osutil

import (
	"errors"
	"io/fs"
	"os"
)

func CheckDirExist(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err != nil {
		// dir not exist
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}
		// other error
		return false, err
	}
	// dir exist
	return info.IsDir(), nil
}

func CheckDirEmpty(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	return len(entries) == 0, nil
}

func RemoveDirIfExist(dir string) error {
	exist, err := CheckDirExist(dir)
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
