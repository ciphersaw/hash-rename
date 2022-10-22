package action

import (
	"fmt"
	"hash-rename/common"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	fileNameChan     chan string
	renameResultChan chan []string
	renameWG         *sync.WaitGroup
	renameCount      int
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
func renameOneFile() error {
	fmt.Printf("Result of renameOneFile:\n")
	// Check if need to rename
	if !argForce {
		needRename := checkIfNeedRename(filepath.Base(argFile))
		if !needRename {
			fmt.Printf("[-] %s has already been renamed with %s value, no need to rename again.\n",
				filepath.Base(argFile), hashFunc.HashName)
			return nil
		}
	}
	// Get file hash
	fileHash, err := hashFunc.GetFileHash(argFile)
	if err != nil {
		fmt.Printf("[-] renameOneFile gets the %s of %s error: %s\n", hashFunc.HashName, argFile, err.Error())
		return err
	}
	if argUppercase {
		fileHash = strings.ToUpper(fileHash)
	}
	// Rename file with its hash value
	fileSuffix := filepath.Ext(argFile)
	fileDir := filepath.Dir(argFile)
	newFile := filepath.Join(fileDir, fileHash+fileSuffix)
	err = os.Rename(argFile, newFile)
	if err != nil {
		fmt.Printf("[-] renameOneFile renames %s to %s error: %s\n", argFile, newFile, err.Error())
		return err
	}
	// Output result
	fmt.Printf("[*] %s --> %s\n", filepath.Base(argFile), filepath.Base(newFile))
	return nil
}

// renameBulkFiles renames the files with specific suffixes in directory.
func renameBulkFiles() error {
	fmt.Printf("Result of renameBulkFiles:\n")
	files, err := os.ReadDir(argDir)
	if err != nil {
		fmt.Printf("[-] renameBulkFiles gets the files in %s error: %s\n", argDir, err.Error())
		return err
	}

	fileNameChan = make(chan string, 64)
	renameResultChan = make(chan []string, 64)
	renameWG = new(sync.WaitGroup)

	// Deal with results
	renameCount = 0
	go func() {
		for result := range renameResultChan {
			renameCount += 1
			fmt.Printf("[%d] %s --> %s\n", renameCount, result[0], result[1])
			renameWG.Done()
		}
	}()

	// Generate goroutines for renaming files
	for c := uint8(0); c < argConcurrency; c++ {
		go func() {
			for fileName := range fileNameChan {
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
				// Check if need to rename
				if !argForce {
					needRename := checkIfNeedRename(fileName)
					if !needRename {
						renameWG.Done()
						continue
					}
				}
				// Get file hash
				oldFile := filepath.Join(argDir, fileName)
				fileHash, err := hashFunc.GetFileHash(oldFile)
				if err != nil {
					fmt.Printf("[-] renameBulkFiles gets the %s of %s error: %s\n", hashFunc.HashName, oldFile, err.Error())
					renameWG.Done()
					continue
				}
				if argUppercase {
					fileHash = strings.ToUpper(fileHash)
				}
				// Rename file with its hash value
				newFile := filepath.Join(argDir, fileHash+fileSuffix)
				err = os.Rename(oldFile, newFile)
				if err != nil {
					fmt.Printf("[-] renameBulkFiles renames %s to %s error: %s\n", oldFile, newFile, err.Error())
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

	// Print the reasons that no files have been renamed
	if renameCount == 0 {
		reasons := fmt.Sprintf("[-] No files have been renamed, and the possible reasons are as follows:\n")
		reasons += fmt.Sprintf(" 1. The suffixes you specify do not match any files.\n")
		reasons += fmt.Sprintf(" 2. The files in %s have already been renamed with %s value, no need to rename again.\n",
			argDir, hashFunc.HashName)
		reasons += fmt.Sprintf(" 3. Errors happen in getting file hash or renaming file with its hash value.\n")
		fmt.Printf(reasons)
	}

	close(fileNameChan)
	close(renameResultChan)
	return nil
}

// checkIfNeedRename checks the file name with regexp if need to rename with its hash value.
func checkIfNeedRename(fileName string) bool {
	realName := common.GetFileNameWithoutExt(fileName)
	if argUppercase {
		if hashFunc.Uppercase.MatchString(realName) {
			return false
		}
	} else {
		if hashFunc.Lowercase.MatchString(realName) {
			return false
		}
	}
	return true
}
