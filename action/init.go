package action

import (
	"errors"
	"fmt"
	"hash-rename/common"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

const VERSION = "v1.0.0"

var (
	argHelp        bool
	argVersion     bool
	argUppercase   bool
	argFile        string
	argDir         string
	argSuffix      string
	argHash        string
	argConcurrency uint8

	err          error
	suffixConfig *SuffixConfig
	hashFunc     *HashFunc
)

// SuffixConfig records the suffix config.
// isSetAll renames all files ignoring suffix.
// isSetNull renames the files without any suffix.
// suffixMap only renames the files with specific suffixes.
type SuffixConfig struct {
	isSetAll  bool
	isSetNull bool
	suffixMap map[string]struct{}
}

// InitArgs initializes and resolves the input arguments.
func InitArgs() {
	flag.BoolVarP(&argHelp, "help", "", false, usageMap["help"])
	flag.BoolVarP(&argVersion, "version", "v", false, usageMap["version"])
	flag.BoolVarP(&argUppercase, "uppercase", "u", false, usageMap["uppercase"])
	flag.StringVarP(&argFile, "file", "f", "", usageMap["file"])
	flag.StringVarP(&argDir, "dir", "d", "", usageMap["dir"])
	flag.StringVarP(&argSuffix, "suffix", "s", "", usageMap["suffix"])
	flag.StringVarP(&argHash, "hash", "h", "md5", usageMap["hash"])
	flag.Uint8VarP(&argConcurrency, "concurrency", "c", 4, usageMap["concurrency"])
	flag.Usage = usage
	flag.Parse()

	if argVersion {
		fmt.Printf("%s\n", VERSION)
		os.Exit(0)
	}

	if argFile != "" {
		if !common.IsPathExist(argFile) {
			os.Exit(1)
		}
	}

	if argDir != "" {
		if !common.IsDirExist(argDir) {
			os.Exit(1)
		}

		suffixConfig, err = setSuffixConfig(argSuffix)
		if err != nil {
			os.Exit(1)
		}

		err = checkConcurrency(argConcurrency)
		if err != nil {
			os.Exit(1)
		}
	}

	if argFile == "" && argDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	if argHash, hashFunc, err = setHashFunc(argHash); err != nil {
		os.Exit(1)
	}
}

// setSuffixConfig sets the suffix config.
func setSuffixConfig(s string) (sc *SuffixConfig, err error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		fmt.Printf("setSuffixConfig gets an empty string.\n")
		return sc, errors.New("empty string")
	}

	sc = new(SuffixConfig)
	s = strings.ToLower(s)
	if s == "all" {
		sc.isSetAll = true
	} else if s == "null" || s == "none" {
		sc.isSetNull = true
	} else {
		sc.suffixMap = make(map[string]struct{})
		// Format the separator and split by comma
		s = strings.ReplaceAll(s, `;`, `,`)
		s = strings.ReplaceAll(s, `|`, `,`)
		s = strings.ReplaceAll(s, ` `, `,`)
		sArr := strings.Split(s, `,`)
		for _, suffix := range sArr {
			suffix = strings.TrimLeft(suffix, `.`)
			if len(suffix) > 0 {
				sc.suffixMap[suffix] = struct{}{}
			}
		}
	}

	return sc, nil
}

// setHashFunc sets the hash function.
func setHashFunc(s string) (key string, hf *HashFunc, err error) {
	key = strings.ToLower(strings.TrimSpace(s))
	if hf, ok := availableHashFunc[key]; ok {
		return key, hf, nil
	} else {
		fmt.Printf("checkHashFunc checks %s hash function is unavailable.\n", key)
		return key, hf, errors.New(key + " hash function is unavailable")
	}
}

// checkConcurrency checks if the goroutine concurrency is valid.
func checkConcurrency(u uint8) (err error) {
	if u >= 1 && u <= 64 {
		return nil
	} else {
		fmt.Printf("checkConcurrency checks %d concurrency is invalid.\n", u)
		return errors.New(strconv.Itoa(int(u)) + " concurrency is invalid")
	}
}
