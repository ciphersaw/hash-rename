package action

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	fileNameChan     = make(chan string, 64)
	renameResultChan = make(chan []string, 64)
	renameWG         = new(sync.WaitGroup)
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

	// Deal with results
	go func() {
		count := 0

		for {
			result := <-renameResultChan
			count += 1
			fmt.Printf("[%d] %s --> %s\n", count, result[0], result[1])
			renameWG.Done()
		}
	}()

	// Generate goroutines for renaming files
	for c := uint8(0); c < argConcurrency; c++ {
		go func() {
			for {
				fileName := <-fileNameChan
				// Check suffix
				fileSuffix := filepath.Ext(fileName)
				if suffixConfig.isSetAll {
					// Do nothing
				} else if suffixConfig.isSetNull {
					if fileSuffix != "" {
						renameWG.Done()
						continue
					}
				} else {
					if _, ok := suffixConfig.suffixMap[strings.TrimLeft(fileSuffix, `.`)]; !ok {
						renameWG.Done()
						continue
					}
				}
				// Get file hash
				oldFile := filepath.Join(argDir, fileName)
				fileHash, err := GetFileHash(oldFile, argHash)
				if err != nil {
					fmt.Printf("renameBulkFiles gets the %s of %s error: %s\n", argHash, oldFile, err.Error())
					renameWG.Done()
					continue
				}
				// Rename file with its hash value
				newFile := filepath.Join(argDir, fileHash+fileSuffix)
				err = os.Rename(oldFile, newFile)
				if err != nil {
					fmt.Printf("renameBulkFiles renames %s to %s error: %s\n", oldFile, newFile, err.Error())
					renameWG.Done()
					continue
				}
				// Output result
				renameResultChan <- []string{filepath.Base(oldFile), filepath.Base(newFile)}
			}
		}()
	}

	// Collect file names for renaming
	renameWG.Add(len(files))
	for _, file := range files {
		fileNameChan <- file.Name()
	}
	renameWG.Wait()
}
