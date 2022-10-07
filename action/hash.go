package action

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

// availableHashFunc records the available hash function currently.
var availableHashFunc = map[string]struct{}{
	"md5":    struct{}{},
	"sha1":   struct{}{},
	"sha256": struct{}{},
}

// GetFileHash gets the hash value of file according to the specific hash function.
func GetFileHash(filePath, hashFunc string) (hash string, err error) {
	switch hashFunc {
	case "md5":
		return GetFileMD5(filePath)
	case "sha1":
		return GetFileSHA1(filePath)
	case "sha256":
		return GetFileSHA256(filePath)
	default:
		return "", errors.New(hashFunc + " hash function is unavailable")
	}
}

// GetFileMD5 calculates the MD5 hash value of file.
func GetFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// GetFileSHA1 calculates the SHA1 hash value of file.
func GetFileSHA1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// GetFileSHA256 calculates the SHA256 hash value of file.
func GetFileSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
