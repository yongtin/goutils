package files

import (
	"fmt"
	"os"
)

// IsFile check if input (filepath) is a file
func IsFile(filepath string) (bool, error) {
	cf, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	switch mode := cf.Mode(); {
	case mode.IsRegular():
		return true, nil
	}
	return false, fmt.Errorf("%s not a regular file", filepath)
}

// IsDir check if input (filepath) is a directory
func IsDir(filepath string) (bool, error) {
	cf, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	switch mode := cf.Mode(); {
	case mode.IsDir():
		return true, nil
	}
	return false, fmt.Errorf("%s not a directory", filepath)
}
