package action

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

// usageMap records the usage of each argument.
var usageMap = map[string]string{
	"help":    "Print the usage of hash-rename.",
	"version": "Print the version of hash-rename.",
	"file":    "Set the file path to be renamed.",
	"dir":     "Set the directory path including files to be renamed.",
	"suffix":  "Set the suffixes of files.",
	"hash":    "Set the hash function for renaming.\nCurrently available for md5, sha1, sha256.",
}

// usage customizes the usage information for hash-rename.
func usage() {
	fmt.Fprintf(os.Stderr, "Usages:\n")
	fmt.Fprintf(os.Stderr, "%s <-f /path/to/file> [-h hash_func]\n", filepath.Base(os.Args[0]))
	fmt.Fprintf(os.Stderr, "%s <-d /path/to/dir -s suffix1,suffix2,...> [-h hash_func]\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}
