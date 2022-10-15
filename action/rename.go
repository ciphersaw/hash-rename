package action

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// StartRenameTask starts the rename task according to initialization arguments.
func StartRenameTask() {
	if argFile != "" {
		renameOneFile()
	}
	if argDir != "" {
		renameBulkFiles()
	}
}

// renameOneFile renames a specific file.
func renameOneFile() {
	fmt.Printf("Result of renameOneFile:\n")
	fileHash, err := GetFileHash(argFile, argHash)
	if err != nil {
		fmt.Printf("renameOneFile gets the %s of %s error: %s\n", argHash, argFile, err.Error())
		return
	}
	fileSuffix := filepath.Ext(argFile)
	fileDir := filepath.Dir(argFile)
	newFile := filepath.Join(fileDir, fileHash+fileSuffix)
	err = os.Rename(argFile, newFile)
	if err != nil {
		fmt.Printf("renameOneFile renames %s to %s error: %s\n", argFile, newFile, err.Error())
		return
	}
	fmt.Printf("[*] %s --> %s\n", filepath.Base(argFile), filepath.Base(newFile))
}

// renameBulkFiles renames the files with specific suffixes in directory.
func renameBulkFiles() {
	fmt.Printf("Result of renameBulkFiles:\n")
	files, err := os.ReadDir(argDir)
	if err != nil {
		fmt.Printf("renameBulkFiles gets the files in %s error: %s\n", argDir, err.Error())
		return
	}

	if suffixConfig.isSetAll {
		count := 0
		for _, file := range files {
			fileName := file.Name()
			oldFile := filepath.Join(argDir, fileName)
			fileHash, err := GetFileHash(oldFile, argHash)
			if err != nil {
				fmt.Printf("renameBulkFiles gets the %s of %s error: %s\n", argHash, oldFile, err.Error())
				continue
			}
			fileSuffix := filepath.Ext(fileName)
			newFile := filepath.Join(argDir, fileHash+fileSuffix)
			err = os.Rename(oldFile, newFile)
			if err != nil {
				fmt.Printf("renameBulkFiles renames %s to %s error: %s\n", oldFile, newFile, err.Error())
				continue
			}
			count += 1
			fmt.Printf("[%d] %s --> %s\n", count, filepath.Base(oldFile), filepath.Base(newFile))
		}
	} else if suffixConfig.isSetNull {
		count := 0
		for _, file := range files {
			fileName := file.Name()
			fileSuffix := filepath.Ext(fileName)
			if fileSuffix != "" {
				continue
			}
			oldFile := filepath.Join(argDir, fileName)
			fileHash, err := GetFileHash(oldFile, argHash)
			if err != nil {
				fmt.Printf("renameBulkFiles gets the %s of %s error: %s\n", argHash, oldFile, err.Error())
				continue
			}
			newFile := filepath.Join(argDir, fileHash)
			err = os.Rename(oldFile, newFile)
			if err != nil {
				fmt.Printf("renameBulkFiles renames %s to %s error: %s\n", oldFile, newFile, err.Error())
				continue
			}
			count += 1
			fmt.Printf("[%d] %s --> %s\n", count, filepath.Base(oldFile), filepath.Base(newFile))
		}
	} else {
		count := 0
		for _, file := range files {
			fileName := file.Name()
			fileSuffix := filepath.Ext(fileName)
			if _, ok := suffixConfig.suffixMap[strings.TrimLeft(fileSuffix, `.`)]; !ok {
				continue
			}
			oldFile := filepath.Join(argDir, fileName)
			fileHash, err := GetFileHash(oldFile, argHash)
			if err != nil {
				fmt.Printf("renameBulkFiles gets the %s of %s error: %s\n", argHash, oldFile, err.Error())
				continue
			}
			newFile := filepath.Join(argDir, fileHash+fileSuffix)
			err = os.Rename(oldFile, newFile)
			if err != nil {
				fmt.Printf("renameBulkFiles renames %s to %s error: %s\n", oldFile, newFile, err.Error())
				continue
			}
			count += 1
			fmt.Printf("[%d] %s --> %s\n", count, filepath.Base(oldFile), filepath.Base(newFile))
		}
	}
}
