package action

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

var (
	usageMap             map[string]string // The usage map of each argument
	usageRenameOneFile   string            // The usage for renaming a specific file
	usageRenameBulkFiles string            // The usage for renaming the files in directory
)

func init() {
	usageMap = map[string]string{
		"help":        "Print the usage of hash-rename.",
		"version":     "Print the version of hash-rename.",
		"file":        "Set the file path to be renamed.",
		"dir":         "Set the directory path including files to be renamed.",
		"suffix":      "Set the suffixes of files.",
		"hash":        "Set the hash function for renaming.\nCurrently available for md5, sha1, sha256.",
		"concurrency": "Set the goroutine concurrency for renaming files.",
		"uppercase":   "Set the uppercase of hash value for renaming.",
		"force":       "Force to rename ignoring file name check.",
	}

	usageRenameOneFile = "%s <-f /path/to/file> [-h hash_func] [-u] [-F]\n"
	usageRenameBulkFiles = "%s <-d /path/to/dir -s suffix1,suffix2,...> [-h hash_func] [-c num] [-u] [-F]\n"
}

// usage customizes the usage information for hash-rename.
func usage() {
	fmt.Fprintf(os.Stderr, "Usages:\n")
	fmt.Fprintf(os.Stderr, usageRenameOneFile, filepath.Base(os.Args[0]))
	fmt.Fprintf(os.Stderr, usageRenameBulkFiles, filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}
