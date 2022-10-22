package action

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"os"
	"regexp"
)

// availableHashFunc records the available hash function currently.
var availableHashFunc = map[string]*HashFunc{
	"md5": {
		HashName:  "md5",
		Lowercase: regexp.MustCompile(`^[0-9a-f]{32}$`),
		Uppercase: regexp.MustCompile(`^[0-9A-F]{32}$`),
	},
	"sha1": {
		HashName:  "sha1",
		Lowercase: regexp.MustCompile(`^[0-9a-f]{40}$`),
		Uppercase: regexp.MustCompile(`^[0-9A-F]{40}$`),
	},
	"sha256": {
		HashName:  "sha256",
		Lowercase: regexp.MustCompile(`^[0-9a-f]{64}$`),
		Uppercase: regexp.MustCompile(`^[0-9A-F]{64}$`),
	},
}

// HashFunc defines the hash function elements during renaming.
type HashFunc struct {
	HashName  string
	Lowercase *regexp.Regexp
	Uppercase *regexp.Regexp
}

// GenHashObj generates the hash object according to the specific hash function name.
func (h *HashFunc) GenHashObj() (hash.Hash, error) {
	switch h.HashName {
	case "md5":
		return md5.New(), nil
	case "sha1":
		return sha1.New(), nil
	case "sha256":
		return sha256.New(), nil
	default:
		return nil, errors.New(h.HashName + " hash function is unavailable")
	}
}

// GetFileHash calculates the hash value of file.
func (h *HashFunc) GetFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash, err := h.GenHashObj()
	if err != nil {
		return "", err
	}

	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
