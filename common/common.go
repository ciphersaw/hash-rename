package common

import (
	"fmt"
	"os"
	"path/filepath"
)

// IsPathExist checks if the path is existed.
func IsPathExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("IsPathExist %s is not existed: %s\n", path, err.Error())
			return false
		} else {
			fmt.Printf("IsPathExist os.Stat(%s) error: %s\n", path, err.Error())
			return false
		}
	} else {
		return true
	}
}

// IsDirExist checks if the path of directory is existed.
func IsDirExist(path string) bool {
	if info, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("IsDirExist %s is not existed: %s\n", path, err.Error())
			return false
		} else {
			fmt.Printf("IsDirExist os.Stat(%s) error: %s\n", path, err.Error())
			return false
		}
	} else if !info.IsDir() {
		fmt.Printf("IsDirExist %s is not a directory\n", path)
		return false
	} else {
		return true
	}
}

// GetFileNameWithoutExt gets the file name without extension.
func GetFileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
